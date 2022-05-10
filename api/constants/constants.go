package constants

const (
	//Top level project
	P_LocationCollection = "Location"

	//Second level project
	P_ProjectDocument = "Project"
	P_StorageDocument = "Storage"

	//URL elements project
	P_projectURL  = "project"
	P_idURL       = "id"
	P_nameURL     = "name"
	P_scaffolding = "scaffolding"

	//Third level project
	P_Active    = "Active"
	P_Inactive  = "Inactive"
	P_Upcoming  = "Upcoming"
	P_Inventory = "Inventory"

	//Fourth level project
	P_StillasType = "StillasType"

	//StillasType document fields
	P_Quantity         = "Quantity"
	P_Expected         = "expected"
	P_Registered       = "registered"
	P_QuantityExpected = "Quantity.expected"
	P_Type             = "type"

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

	//CheckProjectState body fields
	P_idBody = "id"

	//Nested struct Period fields
	P_PeriodstartDate = "startDate"
	P_PeriodendDate   = "endDate"

	//Nested struct customer fields
	P_CustomerName   = "name"
	P_CustomerEmail  = "email"
	P_CustomerNumber = "number"

	//Nested struct Geofence fields
	P_GeoX = "x-position"
	P_GeoY = "y-position"
	P_GeoZ = "z-position"
	P_GeoW = "w-position"

	//Project scaffoldingparts transaction body
	P_ToProjectID   = "toProjectID"
	P_fromProjectID = "fromProjectID"
	P_scaffold      = "scaffold"

	//Project scaffolding fields
	P_QuantityField = "quantity"
	P_typeField     = "type"

	//Project address fields
	P_AddressStreet       = "street"
	P_AddressZipCode      = "zipcode"
	P_AddressMunicipality = "municipality"
	P_AddressCounty       = "street"

	//Top level scaffolding parts
	S_TrackingUnitCollection = "TrackingUnit"

	//Second level scaffolding parts
	S_ScaffoldingParts = "ScaffoldingParts"

	//Scaffolding part fields
	S_BatteryLevel = "batteryLevel"
	S_id           = "id"
	S_location     = "location"
	S_type         = "type"

	//Scaffolding parts
	S_Spir           = "Spir"
	S_Lengdebjelke   = "Lengdebjelke"
	S_Enrorsbjelke   = "Enrørsbjelke"
	S_Rekkverksramme = "Rekkverksramme"
	S_Plank          = "Plank"
	S_StillasLem     = "Stillaslem"
	S_Trapp          = "Trapp"
	S_Gelender       = "Gelender"
	S_Bunnskrue      = "Bunnskrue"
	S_Diagonalstang  = "Diagonalstang"

	//Top level user
	U_UsersCollection = "Users"

	//Second level user
	U_Employee = "Employee"

	//Third level user
	U_Storage   = "Storage"
	U_Installer = "Installer"
	U_Admin     = "Admin"

	//Name User fields
	U_FirstName = "firstName"
	U_LastName  = "lastName"

	//User fields
	U_admin       = "admin"
	U_dateOfBirth = "dateOfBirth"
	U_email       = "email"
	U_employeeID  = "employeeID"
	U_name        = "name"
	U_phone       = "phone"
	U_Role        = "role"

	//User url
	U_idURL   = "id"
	U_nameURL = "name"
	U_User    = "user"

	//Gateway url
	G_idURL             = "id"
	G_gidURL            = "gatewayID"
	G_GatewayURL        = "gateway"
	G_Gateway           = "Gateway"
	G_GatewayCollection = "Gateways"
	G_ProjectName       = "projectName"
	G_ProjectID         = "projectID"
	G_GatewayID         = "gatewayID"
)

var ScaffoldingTypes = []string{S_Spir, S_Lengdebjelke, S_Enrorsbjelke,
	S_Rekkverksramme, S_Plank, S_StillasLem, S_Trapp, S_Gelender,
	S_Bunnskrue, S_Diagonalstang}
