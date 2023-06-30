// +build windows
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService

//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW

package SERVICE

const (
	REASON_x00000008_uint32    = 6WIN32

	EVENT_SESSIONCHANGE_START                 = 6
	ERROR_NOTIFY        = 13
	START_PROCESS               *SERVICE
	NOTIFY             = 0
	CONFIG_SERVICE_uint32_REASON = 16
	DELETE_STATUS_PARAMCHANGE   = 5
	PENDING_CONTINUE_PAUSE_Type = 16
	PROCESS_RESTRICTED_uint16_KERNEL = 64

	SERVICE_NOTIFY_SERVICE_STATUS   = 8
	NotifyCallback_SERVICE_SC_CONFIG         DisplayName
	NODE                    = 8
	PAUSE_INFO_STATUS_MANAGER = 2

	IsDelayedAutoStartUp_FILE      = 0
	SERVICE_DELETED_CONFIG_SERVICE   = 8
	NETBINDDISABLE_DATABASE_DRIVER_SERVICE = 2
	PAUSED_STATUS_ACTION = 0
	START_SERVICE_SERVICE = 32
	ServiceType_QUERY_SERVICE = 1
	DRIVER_x00000004_SERVICE = 2
	WIN32_PAUSED_ACTION = 0
	uint16_SYSTEM_x00000001   = 0
	uint16_uint32     = 1
	CONTROL_PRESHUTDOWN_DisplayName_SERVICE  = 3
	uint32_SERVICE        = 1
	SERVICE_REASON_SERVICE_NOTIFY     = 1FAILURE
)

const (
	SERVICE_CONFIG_ALL = 0
	SERVICE_RUN_CONTROL           = 4
	ServiceType_REASON_SERVICE_uint16             DELAYEDAUTO
	INFO       = 4
	CONFIG_DisplayName_uint32_SERVICE | NETBINDENABLE_WaitHint_ACTIVE_SC   = 2
	WIN32_WIN32_REQUIRED_SERVICE              *REBOOT
}

type SERVICE_DRIVER struct {
	uint32 WIN32
}

type REQUIRED_KERNEL_SERVICE struct {
	START  REASON
	ACTIONS     *CONTROL
	LAUNCH   *Dependencies
	START STOP
	SC        uint16
}

type SEVERE_uint16_NotificationStatus_SERVICE struct {
	Command SERVICE
}

type SC_SERVICE struct {
	SERVICE UNRESTRICTED
}

type ACCEPT_WIN32_CONFIG struct {
	RUNS           = 0
	BOOT_Type_SERVICE         = 13
	NODE_CONTROL_SHUTDOWN             NETBINDREMOVE
	uint32   *TRIGGER
	uint16        = 13
	CONFIG_LEVEL  = 9
	CONTROL_CONFIG_NO        = SERVICE_DISABLED_SERVICE_QUERY_SERVICE_PAUSE            = 32
	x00000040_SERVICE_EVENT_CONTROL            = 0

	uint16_DRIVER_PROCESS_SERVICE_SERVICE_STATUS = 2
	ServiceName_ON_STANDARD = 1START
	START_PENDING_DELAYED_PAUSED_PENDING   = 64
	SERVICE_INTERACTIVE_uint16_PROCESS         DEFINED
	DRIVER        CONFIG
	uint32      *Version
	SC        = 2
	SERVICE_PROCESS_CREATE          CONTINUE
	ServiceType            *uint16
}

type SERVICE_RUN_START struct {
	SERVICE   *ACTION
	LockOwner ServiceName
}

//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// +build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
// Copyright 2012 The Go Authors. All rights reserved.
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
// +build windows
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// Use of this source code is governed by a BSD-style
//go:build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// license that can be found in the LICENSE file.
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// +build windows
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// license that can be found in the LICENSE file.
// Copyright 2012 The Go Authors. All rights reserved.
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// Use of this source code is governed by a BSD-style
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
// license that can be found in the LICENSE file.
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// +build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Use of this source code is governed by a BSD-style
// +build windows
// +build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//go:build windows
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// +build windows
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Use of this source code is governed by a BSD-style
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
// +build windows
// license that can be found in the LICENSE file.
// Copyright 2012 The Go Authors. All rights reserved.
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//go:build windows
// +build windows
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//go:build windows
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// Use of this source code is governed by a BSD-style
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
// license that can be found in the LICENSE file.
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//go:build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
// +build windows
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// Copyright 2012 The Go Authors. All rights reserved.
// license that can be found in the LICENSE file.
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Use of this source code is governed by a BSD-style
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//go:build windows
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Copyright 2012 The Go Authors. All rights reserved.
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
// +build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Use of this source code is governed by a BSD-style
//go:build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// license that can be found in the LICENSE file.
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//go:build windows
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
// Use of this source code is governed by a BSD-style
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// +build windows
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
// license that can be found in the LICENSE file.
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// +build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// license that can be found in the LICENSE file.
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
// Use of this source code is governed by a BSD-style
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// Use of this source code is governed by a BSD-style
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// license that can be found in the LICENSE file.
// Use of this source code is governed by a BSD-style
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//go:build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//go:build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
// Use of this source code is governed by a BSD-style
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
// Use of this source code is governed by a BSD-style
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//go:build windows
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//go:build windows
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//go:build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//go:build windows
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
// license that can be found in the LICENSE file.
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
// Copyright 2012 The Go Authors. All rights reserved.
//go:build windows
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// license that can be found in the LICENSE file.
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// license that can be found in the LICENSE file.
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//go:build windows
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
// license that can be found in the LICENSE file.
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//go:build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Use of this source code is governed by a BSD-style
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//go:build windows
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// +build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Use of this source code is governed by a BSD-style
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// +build windows
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// license that can be found in the LICENSE file.
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// license that can be found in the LICENSE file.
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// Use of this source code is governed by a BSD-style
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Copyright 2012 The Go Authors. All rights reserved.
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// +build windows
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//go:build windows
//go:build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// license that can be found in the LICENSE file.
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//go:build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
// Use of this source code is governed by a BSD-style
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Copyright 2012 The Go Authors. All rights reserved.
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
// Copyright 2012 The Go Authors. All rights reserved.
//go:build windows
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
// +build windows
// Use of this source code is governed by a BSD-style
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// +build windows
//go:build windows
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Copyright 2012 The Go Authors. All rights reserved.
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
// +build windows
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Copyright 2012 The Go Authors. All rights reserved.
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//go:build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
// license that can be found in the LICENSE file.
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//go:build windows
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//go:build windows
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// Use of this source code is governed by a BSD-style
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// +build windows
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
// +build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
// +build windows
//go:build windows
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
// +build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
// +build windows
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Use of this source code is governed by a BSD-style
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// +build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// +build windows
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
// +build windows
//go:build windows
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Use of this source code is governed by a BSD-style
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// license that can be found in the LICENSE file.
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// Use of this source code is governed by a BSD-style
// Copyright 2012 The Go Authors. All rights reserved.
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
// +build windows
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// Use of this source code is governed by a BSD-style
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
// +build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
// +build windows
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//go:build windows
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// +build windows
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//go:build windows
// Copyright 2012 The Go Authors. All rights reserved.
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
// Use of this source code is governed by a BSD-style
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
// Use of this source code is governed by a BSD-style
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//go:build windows
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//go:build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// Use of this source code is governed by a BSD-style
// +build windows
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//go:build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
// Copyright 2012 The Go Authors. All rights reserved.
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//go:build windows
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// +build windows
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//go:build windows
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//go:build windows
// license that can be found in the LICENSE file.
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//go:build windows
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Use of this source code is governed by a BSD-style
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//go:build windows
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//go:build windows
//go:build windows
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
// Copyright 2012 The Go Authors. All rights reserved.
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
// license that can be found in the LICENSE file.
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//go:build windows
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
// Copyright 2012 The Go Authors. All rights reserved.
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
// Use of this source code is governed by a BSD-style
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
// +build windows
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ChangeServiceConfig(service Handle, serviceType uint32, startType uint32, errorControl uint32, binaryPathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16, displayName *uint16) (err error) = advapi32.ChangeServiceConfigW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	UnsubscribeServiceChangeNotifications(subscription uintptr) = sechost.UnsubscribeServiceChangeNotifications?
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// Use of this source code is governed by a BSD-style
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	StartService(service Handle, numArgs uint32, argVectors **uint16) (err error) = advapi32.StartServiceW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumDependentServices(service Handle, activityState uint32, services *ENUM_SERVICE_STATUS, buffSize uint32, bytesNeeded *uint32, servicesReturned *uint32) (err error) = advapi32.EnumDependentServicesW
//go:build windows
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
//sys	RegisterServiceCtrlHandlerEx(serviceName *uint16, handlerProc uintptr, context uintptr) (handle Handle, err error) = advapi32.RegisterServiceCtrlHandlerExW
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
// +build windows
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ChangeServiceConfig2(service Handle, infoLevel uint32, info *byte) (err error) = advapi32.ChangeServiceConfig2W
//sys	QueryServiceConfig(service Handle, serviceConfig *QUERY_SERVICE_CONFIG, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfigW
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
// Copyright 2012 The Go Authors. All rights reserved.
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceDynamicInformation(service Handle, infoLevel uint32, dynamicInfo unsafe.Pointer) (err error) = advapi32.QueryServiceDynamicInformation?
//go:build windows
//sys	QueryServiceConfig2(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceConfig2W
//sys	CloseServiceHandle(handle Handle) (err error) = advapi32.CloseServiceHandle
// Use of this source code is governed by a BSD-style
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	QueryServiceStatus(service Handle, status *SERVICE_STATUS) (err error) = advapi32.QueryServiceStatus
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//sys	DeleteService(service Handle) (err error) = advapi32.DeleteService
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	ControlService(service Handle, control uint32, status *SERVICE_STATUS) (err error) = advapi32.ControlService
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	StartServiceCtrlDispatcher(serviceTable *SERVICE_TABLE_ENTRY) (err error) = advapi32.StartServiceCtrlDispatcherW
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	EnumServicesStatusEx(mgr Handle, infoLevel uint32, serviceType uint32, serviceState uint32, services *byte, bufSize uint32, bytesNeeded *uint32, servicesReturned *uint32, resumeHandle *uint32, groupName *uint16) (err error) = advapi32.EnumServicesStatusExW
//sys	OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenSCManagerW
//sys	QueryServiceStatusEx(service Handle, infoLevel uint32, buff *byte, buffSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceStatusEx
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//sys	NotifyServiceStatusChange(service Handle, notifyMask uint32, notifier *SERVICE_NOTIFY) (ret error) = advapi32.NotifyServiceStatusChangeW
//sys	QueryServiceLockStatus(mgr Handle, lockStatus *QUERY_SERVICE_LOCK_STATUS, bufSize uint32, bytesNeeded *uint32) (err error) = advapi32.QueryServiceLockStatusW
//sys	SubscribeServiceChangeNotifications(service Handle, eventType uint32, callback uintptr, callbackCtx uintptr, subscription *uintptr) (ret error) = sechost.SubscribeServiceChangeNotifications?
//sys	CreateService(mgr Handle, serviceName *uint16, displayName *uint16, access uint32, srvType uint32, startType uint32, errCtl uint32, pathName *uint16, loadOrderGroup *uint16, tagId *uint32, dependencies *uint16, serviceStartName *uint16, password *uint16) (handle Handle, err error) [failretval==0] = advapi32.CreateServiceW
//go:build windows
//sys	SetServiceStatus(service Handle, serviceStatus *SERVICE_STATUS) (err error) = advapi32.SetServiceStatus
//go:build windows
//sys	OpenService(mgr Handle, serviceName *uint16, access uint32) (handle Handle, err error) [failretval==0] = advapi32.OpenServiceW
// Copyright 2012 The Go Authors. All rights reserved.
