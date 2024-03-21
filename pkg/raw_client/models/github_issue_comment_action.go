package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GithubIssueCommentAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The formatter property
    formatter *string
    // The issue property
    issue *string
    // The issueUrl property
    issueUrl *string
    // The owner property
    owner *string
    // The repo property
    repo *string
}
// NewGithubIssueCommentAction instantiates a new GithubIssueCommentAction and sets the default values.
func NewGithubIssueCommentAction()(*GithubIssueCommentAction) {
    m := &GithubIssueCommentAction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGithubIssueCommentActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGithubIssueCommentActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGithubIssueCommentAction(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GithubIssueCommentAction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GithubIssueCommentAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["issue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssue(val)
        }
        return nil
    }
    res["issueUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueUrl(val)
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
    return res
}
// GetFormatter gets the formatter property value. The formatter property
// returns a *string when successful
func (m *GithubIssueCommentAction) GetFormatter()(*string) {
    return m.formatter
}
// GetIssue gets the issue property value. The issue property
// returns a *string when successful
func (m *GithubIssueCommentAction) GetIssue()(*string) {
    return m.issue
}
// GetIssueUrl gets the issueUrl property value. The issueUrl property
// returns a *string when successful
func (m *GithubIssueCommentAction) GetIssueUrl()(*string) {
    return m.issueUrl
}
// GetOwner gets the owner property value. The owner property
// returns a *string when successful
func (m *GithubIssueCommentAction) GetOwner()(*string) {
    return m.owner
}
// GetRepo gets the repo property value. The repo property
// returns a *string when successful
func (m *GithubIssueCommentAction) GetRepo()(*string) {
    return m.repo
}
// Serialize serializes information the current object
func (m *GithubIssueCommentAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formatter", m.GetFormatter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issue", m.GetIssue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issueUrl", m.GetIssueUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GithubIssueCommentAction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormatter sets the formatter property value. The formatter property
func (m *GithubIssueCommentAction) SetFormatter(value *string)() {
    m.formatter = value
}
// SetIssue sets the issue property value. The issue property
func (m *GithubIssueCommentAction) SetIssue(value *string)() {
    m.issue = value
}
// SetIssueUrl sets the issueUrl property value. The issueUrl property
func (m *GithubIssueCommentAction) SetIssueUrl(value *string)() {
    m.issueUrl = value
}
// SetOwner sets the owner property value. The owner property
func (m *GithubIssueCommentAction) SetOwner(value *string)() {
    m.owner = value
}
// SetRepo sets the repo property value. The repo property
func (m *GithubIssueCommentAction) SetRepo(value *string)() {
    m.repo = value
}
type GithubIssueCommentActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormatter()(*string)
    GetIssue()(*string)
    GetIssueUrl()(*string)
    GetOwner()(*string)
    GetRepo()(*string)
    SetFormatter(value *string)()
    SetIssue(value *string)()
    SetIssueUrl(value *string)()
    SetOwner(value *string)()
    SetRepo(value *string)()
}
