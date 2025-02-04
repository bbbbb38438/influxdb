// Libraries
import React, {SFC} from 'react'
import {connect} from 'react-redux'
import {
  NINETEEN_EIGHTY_FOUR,
  ATLANTIS,
  DO_ANDROIDS_DREAM,
  DELOREAN,
  CTHULHU,
  ECTOPLASM,
  T_MAX_400_FILM,
} from '@influxdata/giraffe'

// Components
import {Form, Input, Grid, MultiSelectDropdown} from '@influxdata/clockface'
import AxisAffixes from 'src/timeMachine/components/view_options/AxisAffixes'

// Actions
import {
  setFillColumns,
  setSymbolColumns,
  setYAxisLabel,
  setXAxisLabel,
  setAxisPrefix,
  setAxisSuffix,
  setColorHexes,
  setYDomain,
  setXColumn,
  setYColumn,
} from 'src/timeMachine/actions'

// Utils
import {
  getGroupableColumns,
  getFillColumnsSelection,
  getSymbolColumnsSelection,
  getXColumnSelection,
  getYColumnSelection,
  getNumericColumns,
} from 'src/timeMachine/selectors'

// Types
import {ComponentStatus} from '@influxdata/clockface'
import {AppState} from 'src/types'
import HexColorSchemeDropdown from 'src/shared/components/HexColorSchemeDropdown'
import AutoDomainInput from 'src/shared/components/AutoDomainInput'
import ColumnSelector from 'src/shared/components/ColumnSelector'

const COLOR_SCHEMES = [
  {name: 'Nineteen Eighty Four', colors: NINETEEN_EIGHTY_FOUR},
  {name: 'Atlantis', colors: ATLANTIS},
  {name: 'Do Androids Dream of Electric Sheep?', colors: DO_ANDROIDS_DREAM},
  {name: 'Delorean', colors: DELOREAN},
  {name: 'Cthulhu', colors: CTHULHU},
  {name: 'Ectoplasm', colors: ECTOPLASM},
  {name: 'T-MAX 400 Film', colors: T_MAX_400_FILM},
]

interface StateProps {
  fillColumns: string[]
  symbolColumns: string[]
  availableGroupColumns: string[]
  xColumn: string
  yColumn: string
  numericColumns: string[]
}

interface DispatchProps {
  onSetFillColumns: typeof setFillColumns
  onSetSymbolColumns: typeof setSymbolColumns
  onSetColors: typeof setColorHexes
  onSetYAxisLabel: typeof setYAxisLabel
  onSetXAxisLabel: typeof setXAxisLabel
  onUpdateAxisSuffix: typeof setAxisSuffix
  onUpdateAxisPrefix: typeof setAxisPrefix
  onSetYDomain: typeof setYDomain
  onSetXColumn: typeof setXColumn
  onSetYColumn: typeof setYColumn
}

interface OwnProps {
  xColumn: string
  yColumn: string
  fillColumns: string[]
  symbolColumns: string[]
  xDomain: number[]
  yDomain: number[]
  xAxisLabel: string
  yAxisLabel: string
  xPrefix: string
  xSuffix: string
  yPrefix: string
  ySuffix: string
  colors: string[]
  showNoteWhenEmpty: boolean
}

type Props = OwnProps & DispatchProps & StateProps

const ScatterOptions: SFC<Props> = props => {
  const {
    fillColumns,
    symbolColumns,
    availableGroupColumns,
    yAxisLabel,
    xAxisLabel,
    onSetFillColumns,
    onSetSymbolColumns,
    colors,
    onSetColors,
    onSetYAxisLabel,
    onSetXAxisLabel,
    yPrefix,
    ySuffix,
    onUpdateAxisSuffix,
    onUpdateAxisPrefix,
    yDomain,
    onSetYDomain,
    xColumn,
    yColumn,
    numericColumns,
    onSetXColumn,
    onSetYColumn,
  } = props

  const groupDropdownStatus = availableGroupColumns.length
    ? ComponentStatus.Default
    : ComponentStatus.Disabled

  const handleFillColumnSelect = (column: string): void => {
    let updatedFillColumns

    if (fillColumns.includes(column)) {
      updatedFillColumns = fillColumns.filter(col => col !== column)
    } else {
      updatedFillColumns = [...fillColumns, column]
    }

    onSetFillColumns(updatedFillColumns)
  }

  const handleSymbolColumnSelect = (column: string): void => {
    let updatedSymbolColumns

    if (symbolColumns.includes(column)) {
      updatedSymbolColumns = symbolColumns.filter(col => col !== column)
    } else {
      updatedSymbolColumns = [...symbolColumns, column]
    }

    onSetSymbolColumns(updatedSymbolColumns)
  }

  return (
    <Grid.Column>
      <h4 className="view-options--header">Customize Scatter Plot</h4>
      <h5 className="view-options--header">Data</h5>

      <Form.Element label="Symbol Column">
        <MultiSelectDropdown
          options={availableGroupColumns}
          selectedOptions={symbolColumns}
          onSelect={handleSymbolColumnSelect}
          buttonStatus={groupDropdownStatus}
        />
      </Form.Element>
      <Form.Element label="Fill Column">
        <MultiSelectDropdown
          options={availableGroupColumns}
          selectedOptions={fillColumns}
          onSelect={handleFillColumnSelect}
          buttonStatus={groupDropdownStatus}
        />
      </Form.Element>
      <ColumnSelector
        selectedColumn={xColumn}
        onSelectColumn={onSetXColumn}
        availableColumns={numericColumns}
        axisName="x"
      />
      <ColumnSelector
        selectedColumn={yColumn}
        onSelectColumn={onSetYColumn}
        availableColumns={numericColumns}
        axisName="y"
      />
      <h5 className="view-options--header">Options</h5>
      <Form.Element label="Color Scheme">
        <HexColorSchemeDropdown
          colorSchemes={COLOR_SCHEMES}
          selectedColorScheme={colors}
          onSelectColorScheme={onSetColors}
        />
      </Form.Element>
      <h5 className="view-options--header">X Axis</h5>
      <Form.Element label="X Axis Label">
        <Input
          value={xAxisLabel}
          onChange={e => onSetXAxisLabel(e.target.value)}
        />
      </Form.Element>
      <h5 className="view-options--header">Y Axis</h5>
      <Form.Element label="Y Axis Label">
        <Input
          value={yAxisLabel}
          onChange={e => onSetYAxisLabel(e.target.value)}
        />
      </Form.Element>
      <Grid.Row>
        <AxisAffixes
          prefix={yPrefix}
          suffix={ySuffix}
          axisName="y"
          onUpdateAxisPrefix={prefix => onUpdateAxisPrefix(prefix, 'y')}
          onUpdateAxisSuffix={suffix => onUpdateAxisSuffix(suffix, 'y')}
        />
      </Grid.Row>
      <AutoDomainInput
        domain={yDomain as [number, number]}
        onSetDomain={onSetYDomain}
        label="Y Axis Domain"
      />
    </Grid.Column>
  )
}

const mstp = (state: AppState): StateProps => {
  const availableGroupColumns = getGroupableColumns(state)
  const fillColumns = getFillColumnsSelection(state)
  const symbolColumns = getSymbolColumnsSelection(state)
  const xColumn = getXColumnSelection(state)
  const yColumn = getYColumnSelection(state)
  const numericColumns = getNumericColumns(state)

  return {
    availableGroupColumns,
    fillColumns,
    symbolColumns,
    xColumn,
    yColumn,
    numericColumns,
  }
}

const mdtp = {
  onSetFillColumns: setFillColumns,
  onSetSymbolColumns: setSymbolColumns,
  onSetColors: setColorHexes,
  onSetYAxisLabel: setYAxisLabel,
  onSetXAxisLabel: setXAxisLabel,
  onUpdateAxisPrefix: setAxisPrefix,
  onUpdateAxisSuffix: setAxisSuffix,
  onSetYDomain: setYDomain,
  onSetXColumn: setXColumn,
  onSetYColumn: setYColumn,
}

export default connect<StateProps, DispatchProps, OwnProps>(
  mstp,
  mdtp
)(ScatterOptions)
