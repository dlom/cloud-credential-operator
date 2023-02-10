package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartLineFormat 
type WorkbookChartLineFormat struct {
    Entity
    // HTML color code representing the color of lines in the chart.
    color *string
}
// NewWorkbookChartLineFormat instantiates a new workbookChartLineFormat and sets the default values.
func NewWorkbookChartLineFormat()(*WorkbookChartLineFormat) {
    m := &WorkbookChartLineFormat{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.workbookChartLineFormat";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkbookChartLineFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartLineFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartLineFormat(), nil
}
// GetColor gets the color property value. HTML color code representing the color of lines in the chart.
func (m *WorkbookChartLineFormat) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartLineFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetColor)
    return res
}
// Serialize serializes information the current object
func (m *WorkbookChartLineFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. HTML color code representing the color of lines in the chart.
func (m *WorkbookChartLineFormat) SetColor(value *string)() {
    m.color = value
}