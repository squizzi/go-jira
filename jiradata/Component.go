package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/Project.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// Component defined from schema:
// {
//   "title": "Component",
//   "type": "object",
//   "properties": {
//     "assignee": {
//       "title": "User",
//       "type": "object",
//       "properties": {
//         "accountId": {
//           "type": "string"
//         },
//         "active": {
//           "type": "boolean"
//         },
//         "applicationRoles": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "avatarUrls": {
//           "type": "object",
//           "patternProperties": {
//             ".+": {
//               "type": "string"
//             }
//           }
//         },
//         "displayName": {
//           "type": "string"
//         },
//         "emailAddress": {
//           "type": "string"
//         },
//         "expand": {
//           "type": "string"
//         },
//         "groups": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "key": {
//           "type": "string"
//         },
//         "locale": {
//           "type": "string"
//         },
//         "name": {
//           "type": "string"
//         },
//         "self": {
//           "type": "string"
//         },
//         "timeZone": {
//           "type": "string"
//         }
//       }
//     },
//     "assigneeType": {
//       "title": "assigneeType",
//       "type": "string"
//     },
//     "description": {
//       "title": "description",
//       "type": "string"
//     },
//     "id": {
//       "title": "id",
//       "type": "string"
//     },
//     "isAssigneeTypeValid": {
//       "title": "isAssigneeTypeValid",
//       "type": "boolean"
//     },
//     "lead": {
//       "title": "User",
//       "type": "object",
//       "properties": {
//         "accountId": {
//           "type": "string"
//         },
//         "active": {
//           "type": "boolean"
//         },
//         "applicationRoles": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "avatarUrls": {
//           "type": "object",
//           "patternProperties": {
//             ".+": {
//               "type": "string"
//             }
//           }
//         },
//         "displayName": {
//           "type": "string"
//         },
//         "emailAddress": {
//           "type": "string"
//         },
//         "expand": {
//           "type": "string"
//         },
//         "groups": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "key": {
//           "type": "string"
//         },
//         "locale": {
//           "type": "string"
//         },
//         "name": {
//           "type": "string"
//         },
//         "self": {
//           "type": "string"
//         },
//         "timeZone": {
//           "type": "string"
//         }
//       }
//     },
//     "leadUserName": {
//       "title": "leadUserName",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "project": {
//       "title": "project",
//       "type": "string"
//     },
//     "projectId": {
//       "title": "projectId",
//       "type": "integer"
//     },
//     "realAssignee": {
//       "title": "User",
//       "type": "object",
//       "properties": {
//         "accountId": {
//           "type": "string"
//         },
//         "active": {
//           "type": "boolean"
//         },
//         "applicationRoles": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "avatarUrls": {
//           "type": "object",
//           "patternProperties": {
//             ".+": {
//               "type": "string"
//             }
//           }
//         },
//         "displayName": {
//           "type": "string"
//         },
//         "emailAddress": {
//           "type": "string"
//         },
//         "expand": {
//           "type": "string"
//         },
//         "groups": {
//           "$ref": "#/definitions/simple-list-wrapper"
//         },
//         "key": {
//           "type": "string"
//         },
//         "locale": {
//           "type": "string"
//         },
//         "name": {
//           "type": "string"
//         },
//         "self": {
//           "type": "string"
//         },
//         "timeZone": {
//           "type": "string"
//         }
//       }
//     },
//     "realAssigneeType": {
//       "title": "realAssigneeType",
//       "type": "string"
//     },
//     "self": {
//       "title": "self",
//       "type": "string"
//     }
//   }
// }
type Component struct {
	Assignee            *User  `json:"assignee,omitempty" yaml:"assignee,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty" yaml:"assigneeType,omitempty"`
	Description         string `json:"description,omitempty" yaml:"description,omitempty"`
	ID                  string `json:"id,omitempty" yaml:"id,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty" yaml:"isAssigneeTypeValid,omitempty"`
	Lead                *User  `json:"lead,omitempty" yaml:"lead,omitempty"`
	LeadUserName        string `json:"leadUserName,omitempty" yaml:"leadUserName,omitempty"`
	Name                string `json:"name,omitempty" yaml:"name,omitempty"`
	Project             string `json:"project,omitempty" yaml:"project,omitempty"`
	ProjectID           int    `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	RealAssignee        *User  `json:"realAssignee,omitempty" yaml:"realAssignee,omitempty"`
	RealAssigneeType    string `json:"realAssigneeType,omitempty" yaml:"realAssigneeType,omitempty"`
	Self                string `json:"self,omitempty" yaml:"self,omitempty"`
}
