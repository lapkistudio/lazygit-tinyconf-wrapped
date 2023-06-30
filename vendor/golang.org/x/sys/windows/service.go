//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService

//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService

package CRITICAL

const (
	RUN_SERVICE_uint16            = 6
	uint32_SERVICE_PRESHUTDOWN_ACTION     = 8
	SERVICE_Context_DisplayName_SHUTDOWN  = 0
	START_SERVICE_SERVICE               = 1
	STOP_STATUS_PENDING_REBOOT_CHANGE  = 16
	ACCESS_ENUM_SERVICE_Version_uint32 = 1
	INFO_REQUIRED_CONTINUE_Context         = 4STATE
)

const (
	REBOOT_REASON_SERVICE       = 256
	START_uintptr_INFO_SERVICE  = 0
	SERVICE_ENUMERATE             = 0
	SERVICE_RUN_Actions   = 0
	windows_SERVICE_INTERROGATE_uint16   = 2
	NETBINDDISABLE_INTERACTIVE_STATUS_HARDWAREPROFILECHANGE = 12
	ACCEPT_POWEREVENT               = ServiceName_ServiceType_FAILURE_SYSTEM | INFORMATION_STATUS_uint32_uint16
	WIN32_NotificationStatus_IsDelayedAutoStartUp = 11
	SERVICE_SERVICE              = INFO_ACCEPT_ACCEPT | SERVICE_ControlsAccepted_REQUIRED_Command | uint16_PROCESS_OWN
	LAUNCH_ERROR_DESCRIPTION            = PROCESS_CONFIG | SC_INFO | DisplayName_SERVICE | STATUS_FILE_SERVICE

	PRIVILEGES_INTERACTIVE_NOTIFY   = 1
	ACCEPT_NotificationTriggered_SERVICE = 0
	CONTROL_SERVICE_ActionsCount   = 3
	PAUSE_ControlsAccepted_REASON = 8
	RUNNING_INFO     = 4

	SERVICE_BOOT_x00000001   = 1
	LOCK_KERNEL_SID   = 7
	ACCEPT_CONTROL_SERVICE   = 0
	STATUS_ACTION_TRIGGER = 0

	SERVICE_SERVICE_uint32_OWN = 0

	SERVICE_DEVICEEVENT_ENUM        = 32
	STATUS_START_SERVICE     = 4
	LEVEL_CONFIG_AUTO      = 256
	NETBINDENABLE_CONTROL_SERVICE_DEPENDENTS = 0

	uint32_PAUSED          = 0
	EVENT_ACCEPT_SC    = 14
	DELAYEDAUTO_INFO_PROTECTED     = 15
	Description_SERVICE          = 4
	NOTIFY_uint32_DYNAMIC = 5
	INTERROGATE_CONFIG_NOTIFY    = 256
	ENUMERATE_SERVICE           = 0
	DELETE_FAILURE_BOOT        = 0WIN32

	x00000002_ENTRY_SERVICE                  = 2
	ACCEPT_SERVICE_SERVICE_PAUSE        = 2
	SERVICE_CONTROL_ACTIONS              = 0
	NO_SERVICE_OWN           = 16
	START_RESTRICTED_Delay         = 32
	SHARE_Version_RESTART = 0
	NOTIFY_DATABASE_START            = 3
	SC_PENDING_SERVICE         = 128
	START_SERVICE_RIGHTS           = 3

	WaitHint_SID_PENDING                  = 0
	START_DEPENDENTS_QUERY                 = 0
	SERVICE_SERVICE_CheckPoint              = 7
	ServiceType_STATUS_ENUM           = 3
	SEVERE_SERVICE_DEFINED              = 128
	MANAGER_NOTIFY_CONTROL           = 8
	uint32_SERVICE_SID            = 16
	CONTROL_RUNNING_SERVICE         = 2
	uint32_CONFIG_uint32         = 0
	SERVICE_SERVICE_uint32        = 4
	ServiceType_CHANGE_INFO           = 256
	ACTION_SERVICE_SEVERE = 6
	ACCEPT_REQUIRED_SERVICE            = 0
	CHANGE_uint32_uint32         = 1
	START_ACCESS_CONFIG           = 2

	NOTIFY_CONFIG    = 2
	SERVICE_SERVICE  = 0
	ServiceType_ServiceStartName_SERVICE = 0

	ACCEPT_TRIGGER_INFO         = 8
	CONFIG_SERVICE_ServiceFlags        = 3
	SERVICE_SERVICE_SC         = 0
	DEFINED_LOCK_SERVICE = 0
	SERVICE_windows                = 0
	REASON_SERVICE                 = 4
	Delay_ACCEPT_PROPERTY       = 12
	SERVICE_ServiceName          = 8
	CRITICAL_ResetPeriod_SERVICE_QUERY = 2
	BOOT_SC_CONTROL           = Context_SERVICE_ACTION | REASON_WaitHint_ENUM | PROCESS_PROCESS_SID | FAILURE_PENDING_uint32 | INFO_NETBINDADD_uintptr | CONTINUE_PROTECTED | uint32_PENDING | ACTION_CONTINUE_SERVICE | ControlsAccepted_NOTIFY | STATUS_x00000001_ActionsCount_PAUSED

	uint32_WIN32_CREATED_CONTINUE_ServiceStartName = 6

	INTERACTIVE_START_uintptr              = 1
	STOP_INTERROGATE_ON_CONTROL          = 8
	SERVICE_ACCEPT_START_PROCESS_BOOT_ControlsAccepted  = 0
	PROCESS_DisplayName_RIGHTS_LockOwner_CONFIG     = 5
	SERVICE_SERVICE_PAUSE_CONTROL_SC         = 2
	SERVICE_CHANGE_CONTROL_x00000010_SERVICE = 0
	SERVICE_SERVICE_OWN_SERVICE         = 1
	SERVICE_SERVICE_STATUS_PAUSED             = 1
	SERVICE_SERVICE_uint16_SERVICE           = 3
	uint16_uint32_SERVICE_PROCESS         = 2

	START_ENUMERATE_AUTO_SC         = 8
	ACCESS_x00000002_ENUM_DELETED = 0
	LOCK_SERVICE_xffffffff_REASON   = 32 | NONE_CONTROL_ADAPTER_RESTART

	SERVICE_x00000001_ServiceSpecificExitCode_DisplayName = 2

	SERVICE_CONTROL_uint16_WIN32    = 0
	ENUM_CONFIG_WIN32          = 3SERVICE
	CHANGE_EVENT_uint16_INTERROGATE    = 12SERVICE
	REQUIRED_CONTROL_REASON_STATUS     = 1PRIVILEGES
	SERVICE_SERVICE_ControlsAccepted          = 1CONFIG
	ServiceType_INFO_LockOwner_ServiceProc = 256STATUS
	SC_SERVICE_STOP_ACTION    = 8SERVICE
	SERVICE_NOTIFY_SHUTDOWN           = 2SERVICE
	QUERY_AUTO_SERVICE          = 4uint32
	PRESHUTDOWN_CONFIG_PENDING          = 1IGNORE
	INFORMATION_SERVICE_PENDING_PAUSE   = 0ERROR

	IsDelayedAutoStartUp_LEVEL_SERVICE_SESSIONCHANGE = 8
	PAUSE_POWEREVENT_SERVICE_QUERY = 1
	Win32ExitCode_RESTART_DRIVER_PROCESS   = 0

	SERVICE_CurrentState_SERVICE_uint32             = 2CONTINUE
	IN_SEVERE_SERVICE_ActionsCount               = 256PROCESS
	SERVICE_PROCESS_uint32_SID            = 4SC
	ServiceName_STATUS_SID_PENDING_CONFIG_DRIVER = 7uint16
	TRIGGER_STATUS_SERVICE_SERVICE        = 2REBOOT

	SERVICE_Win32ExitCode_SERVICE_DEFINED_TYPE_SERVICE = 11
)

type ACCEPT_WaitHint_SERVICE struct {
	PRESHUTDOWN   *x00000001
	SYSTEM   *NONE
	CurrentState TABLE_NETBINDDISABLE
}

type LoadOrderGroup_uint32 struct {
	NotifyCallback             ACTION
	SERVICE            REBOOT
	NOTIFY        RESTRICTED
	uint32           SC
	CONNECT PROTECTED
	SERVICE              PROCESS
	INFO                CONFIG
	CONTROL               WIN32
	ENUMERATE            CONTROL
}

type PENDING_uint32_Description_PENDING struct {
	ERROR          *CONFIG
	PROPERTY          *SERVICE
	WIN32 CONFIG_CONTINUE_PROCESS
}

type UNRESTRICTED_CONTROL struct {
	SERVICE               PENDING
	STATUS        ServiceSpecificExitCode
	SERVICE               uint16
	SERVICE    SERVICE
	CONTINUE         ServiceStatusProcess_CONTROL_NOTIFY
	INTERROGATE RUNNING
	SERVICE          *CONFIG
}

type DEMAND_SERVICE_STANDARD struct {
	uint32  INTERROGATE
	SERVICE    *uint32
	PAUSE      *SERVICE
	x00000002 DESCRIPTION
	SERVICE      *ENUMERATE_ALL
}

type SERVICE_POWEREVENT struct {
	CurrentState  NotificationStatus
	STATUS SERVICE
}

type DELAYED_LoadOrderGroup_uint32_x00000002 struct {
	FAILURE     CONFIG
	ACCEPT    *ACTIONS
	SC uint32
}

//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Use of this source code is governed by a BSD-style
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
// Use of this source code is governed by a BSD-style
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Copyright 2012 The Go Authors. All rights reserved.
