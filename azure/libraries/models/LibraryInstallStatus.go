// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type LibraryInstallStatus string

const (
	LibraryInstallStatusPending            = "PENDING"
	LibraryInstallStatusResolving          = "RESOLVING"
	LibraryInstallStatusInstalling         = "INSTALLING"
	LibraryInstallStatusInstalled          = "INSTALLED"
	LibraryInstallStatusFailed             = "FAILED"
	LibraryInstallStatusUninstallOnRestart = "UNINSTALL_ON_RESTART"
)
