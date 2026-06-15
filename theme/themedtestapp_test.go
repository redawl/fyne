package theme_test

import (
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type themedApp struct {
	primaryColor string
	theme        fyne.Theme
	variant      fyne.ThemeVariant
}

func (a *themedApp) CloudProvider() fyne.CloudProvider {
	return nil
}

func (a *themedApp) BuildType() fyne.BuildType {
	return fyne.BuildStandard
}

func (a *themedApp) NewWindow(_ string) fyne.Window {
	return nil
}

func (a *themedApp) OpenURL(_ *url.URL) error {
	return nil
}

func (a *themedApp) Icon() fyne.Resource {
	return nil
}

func (a *themedApp) SetIcon(fyne.Resource) {
}

func (a *themedApp) Run() {
}

func (a *themedApp) Quit() {
}

func (a *themedApp) Driver() fyne.Driver {
	return nil
}

func (a *themedApp) UniqueID() string {
	return ""
}

func (a *themedApp) SendNotification(_ *fyne.Notification) {
}

func (a *themedApp) ScheduleNotification(_ *fyne.Notification, _ time.Time) (*fyne.ScheduledNotification, error) {
	return nil, nil
}

func (a *themedApp) CancelScheduledNotification(_ string) error {
	return nil
}

func (a *themedApp) Settings() fyne.Settings {
	return a
}

func (a *themedApp) Storage() fyne.Storage {
	return nil
}

func (a *themedApp) Preferences() fyne.Preferences {
	return nil
}

func (a *themedApp) Lifecycle() fyne.Lifecycle {
	return nil
}

func (a *themedApp) Metadata() fyne.AppMetadata {
	return fyne.AppMetadata{}
}

func (a *themedApp) PrimaryColor() string {
	if a.primaryColor != "" {
		return a.primaryColor
	}

	return theme.ColorBlue
}

func (a *themedApp) Theme() fyne.Theme {
	return a.theme
}

func (a *themedApp) SetTheme(t fyne.Theme) {
	a.theme = t
}

func (a *themedApp) ThemeVariant() fyne.ThemeVariant {
	return a.variant // The null value is theme.VariantDark
}

func (a *themedApp) SetCloudProvider(fyne.CloudProvider) {
}

func (a *themedApp) Scale() float32 {
	return 1.0
}

func (a *themedApp) ShowAnimations() bool {
	return true
}

func (a *themedApp) AddChangeListener(chan fyne.Settings) {
}

func (a *themedApp) AddListener(func(fyne.Settings)) {
}

func (a *themedApp) Cache() fyne.Cache {
	return nil
}

func (a *themedApp) Clipboard() fyne.Clipboard {
	return nil
}
