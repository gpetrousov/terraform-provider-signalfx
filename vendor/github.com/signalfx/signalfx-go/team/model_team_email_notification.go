/*
 * Teams API
 *
 *  ## Overview An API for creating, retrieving, updating, and deleting teams ## Authentication To authenticate with SignalFx, the following operations require a session  token associated with a SignalFx user that has administrative privileges:<br>   * Create a team - **POST** `/team`   * Update a team - **PUT** `/team/{id}`   * Delete a team - **DELETE** `/team/{id}`   * Update team members - **PUT** `/team/{id}/members`  You can authenticate the following operations with either an org token or a session token. The session token  doesn't need to be associated with a SignalFx user that has administrative privileges:<br>   * Retrieve teams using a query - **GET** `/team`   * Retrieve a team using its ID - **GET** `/team/{id}`
 *
 * API version: 3.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package team

// Properties for a notification sent to the team via email
type TeamEmailNotification struct {
	// Tells SignalFx which system it should use to send the notification. For an TeamEmail notification, this is always \"TeamEmail\".
	Type string `json:"type"`
	// The SignalFx-assigned ID of the team that should receive the notification.
	Team string `json:"team,omitempty"`
}
