/*
 * Teams API
 *
 *  ## Overview An API for creating, retrieving, updating, and deleting teams ## Authentication To authenticate with SignalFx, the following operations require a session  token associated with a SignalFx user that has administrative privileges:<br>   * Create a team - **POST** `/team`   * Update a team - **PUT** `/team/{id}`   * Delete a team - **DELETE** `/team/{id}`   * Update team members - **PUT** `/team/{id}/members`  You can authenticate the following operations with either an org token or a session token. The session token  doesn't need to be associated with a SignalFx user that has administrative privileges:<br>   * Retrieve teams using a query - **GET** `/team`   * Retrieve a team using its ID - **GET** `/team/{id}`
 *
 * API version: 3.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package team

// Notification properties for an alert sent via xMatters
type XMattersNotification struct {
	// Tells SignalFx which external system it should use to send the notification. For an xMatters notification, this is always  \"XMatters\" (with a capital \"X\").
	Type string `json:"type"`
	// xMatters-supplied credential ID that SignalFx uses to authenticate the notification with the xMatters system. Get this value from your xMatters account settings.
	CredentialId string `json:"credentialId"`
}
