// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EntityStatusCode string

// Enum values for EntityStatusCode
const (
	EntityStatusCodeImpaired   EntityStatusCode = "IMPAIRED"
	EntityStatusCodeUnimpaired EntityStatusCode = "UNIMPAIRED"
	EntityStatusCodeUnknown    EntityStatusCode = "UNKNOWN"
)

type EventAggregateField string

// Enum values for EventAggregateField
const (
	EventAggregateFieldEventtypecategory EventAggregateField = "eventTypeCategory"
)

type EventScopeCode string

// Enum values for EventScopeCode
const (
	EventScopeCodePublic           EventScopeCode = "PUBLIC"
	EventScopeCodeAccount_specific EventScopeCode = "ACCOUNT_SPECIFIC"
	EventScopeCodeNone             EventScopeCode = "NONE"
)

type EventStatusCode string

// Enum values for EventStatusCode
const (
	EventStatusCodeOpen     EventStatusCode = "open"
	EventStatusCodeClosed   EventStatusCode = "closed"
	EventStatusCodeUpcoming EventStatusCode = "upcoming"
)

type EventTypeCategory string

// Enum values for EventTypeCategory
const (
	EventTypeCategoryIssue                EventTypeCategory = "issue"
	EventTypeCategoryAccount_notification EventTypeCategory = "accountNotification"
	EventTypeCategoryScheduled_change     EventTypeCategory = "scheduledChange"
	EventTypeCategoryInvestigation        EventTypeCategory = "investigation"
)
