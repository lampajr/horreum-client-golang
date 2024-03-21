package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GithubIssueCreateAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The formatter property
    formatter *string
    // The owner property
    owner *string
    // The repo property
    repo *string
    // The title property
    title *string
}
// NewGithubIssueCreateAction instantiates a new GithubIssueCreateAction and sets the default values.
func NewGithubIssueCreateAction()(*GithubIssueCreateAction) {
    m := &GithubIssueCreateAction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGithubIssueCreateActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGithubIssueCreateActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGithubIssueCreateAction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GithubIssueCreateAction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GithubIssueCreateAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["formatter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormatter(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepo(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFormatter gets the formatter property value. The formatter property
// returns a *string when successful
func (m *GithubIssueCreateAction) GetFormatter()(*string) {
    return m.formatter
}
// GetOwner gets the owner property value. The owner property
// returns a *string when successful
func (m *GithubIssueCreateAction) GetOwner()(*string) {
    return m.owner
}
// GetRepo gets the repo property value. The repo property
// returns a *string when successful
func (m *GithubIssueCreateAction) GetRepo()(*string) {
    return m.repo
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *GithubIssueCreateAction) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *GithubIssueCreateAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formatter", m.GetFormatter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repo", m.GetRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GithubIssueCreateAction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormatter sets the formatter property value. The formatter property
func (m *GithubIssueCreateAction) SetFormatter(value *string)() {
    m.formatter = value
}
// SetOwner sets the owner property value. The owner property
func (m *GithubIssueCreateAction) SetOwner(value *string)() {
    m.owner = value
}
// SetRepo sets the repo property value. The repo property
func (m *GithubIssueCreateAction) SetRepo(value *string)() {
    m.repo = value
}
// SetTitle sets the title property value. The title property
func (m *GithubIssueCreateAction) SetTitle(value *string)() {
    m.title = value
}
type GithubIssueCreateActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormatter()(*string)
    GetOwner()(*string)
    GetRepo()(*string)
    GetTitle()(*string)
    SetFormatter(value *string)()
    SetOwner(value *string)()
    SetRepo(value *string)()
    SetTitle(value *string)()
}
