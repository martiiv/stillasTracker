package constants

const (
	//Top level project
	P_LocationCollection = "Location"

	//Second level project
	P_ProjectDocument = "Project"
	P_StorageDocument = "Storage"

	//Third level project
	P_Active    = "Active"
	P_Inactive  = "Inactive"
	P_Inventory = "Inventory"

	//Project document fields
	P_Address     = "address"
	P_Customer    = "customer"
	P_Geofence    = "geofence"
	P_Latitude    = "latitude"
	P_Longitude   = "longitude"
	P_Period      = "period"
	P_ProjectId   = "projectID"
	P_ProjectName = "projectName"
	P_Scaffolding = "scaffolding"
	P_Size        = "size"
	P_State       = "state"

	//Top level scaffolding parts
	S_TrackingUnitCollection = "TrackingUnit"

	//Second level scaffolding parts
	S_ScaffoldingParts = "ScaffoldingParts"

	//Scaffolding part fields
	S_BatteryLevel = "batteryLevel"
	S_id           = "id"
	S_location     = "location"
	S_type         = "type"

	//Top level user
	U_UsersCollection = "Users"

	//Second level user
	U_Employee = "Employee"

	//Third level user
	U_Storage   = "Storage"
	U_Installer = "Installer"

	//User fields
	U_admin       = "admin"
	U_dateOfBirth = "dateOfBirth"
	U_email       = "email"
	U_employeeID  = "employeeID"
	U_name        = "name"
	U_phone       = "phone"
	U_Role        = "role"
)
