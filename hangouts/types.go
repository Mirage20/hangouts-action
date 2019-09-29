package hangouts

// ActionParameter: List of string parameters to supply when the action
// method is invoked.
// For example, consider three snooze buttons: snooze now, snooze 1
// day,
// snooze next week. You might use action method = snooze(), passing
// the
// snooze type and snooze time in the list of string parameters.
type ActionParameter struct {
	// Key: The name of the parameter for the action script.
	Key string `json:"key,omitempty"`

	// Value: The value of the parameter.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// ActionResponse: Parameters that a bot can use to configure how it's
// response is posted.
type ActionResponse struct {
	// Type: The type of bot response.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default type; will be handled as NEW_MESSAGE.
	//   "NEW_MESSAGE" - Post as a new message in the topic.
	//   "UPDATE_MESSAGE" - Update the bot's own message. (Only after
	// CARD_CLICKED events.)
	//   "REQUEST_CONFIG" - Privately ask the user for additional auth or
	// config.
	Type string `json:"type,omitempty"`

	// Url: URL for users to auth or config. (Only for REQUEST_CONFIG
	// response types.)
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Annotation: Annotations associated with the plain-text body of the
// message.
//
// Example plain-text message body:
// ```
// Hello @FooBot how are you!"
// ```
//
// The corresponding annotations metadata:
// ```
// "annotations":[{
//   "type":"USER_MENTION",
//   "startIndex":6,
//   "length":7,
//   "userMention": {
//     "user": {
//       "name":"users/107946847022116401880",
//       "displayName":"FooBot",
//       "avatarUrl":"https://goo.gl/aeDtrS",
//       "type":"BOT"
//     },
//     "type":"MENTION"
//    }
// }]
// ```
type Annotation struct {
	// Length: Length of the substring in the plain-text message body this
	// annotation
	// corresponds to.
	Length int64 `json:"length,omitempty"`

	// StartIndex: Start index (0-based, inclusive) in the plain-text
	// message body this
	// annotation corresponds to.
	StartIndex int64 `json:"startIndex,omitempty"`

	// Type: The type of this annotation.
	//
	// Possible values:
	//   "ANNOTATION_TYPE_UNSPECIFIED" - Default value for the enum. DO NOT
	// USE.
	//   "USER_MENTION" - A user is mentioned.
	Type string `json:"type,omitempty"`

	// UserMention: The metadata of user mention.
	UserMention *UserMentionMetadata `json:"userMention,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Length") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Length") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Button: A button. Can be a text button or an image button.
type Button struct {
	// ImageButton: A button with image and onclick action.
	ImageButton *ImageButton `json:"imageButton,omitempty"`

	// TextButton: A button with text and onclick action.
	TextButton *TextButton `json:"textButton,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageButton") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageButton") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Card: A card is a UI element that can contain UI widgets such as
// texts, images.
type Card struct {
	// CardActions: The actions of this card.
	CardActions []*CardAction `json:"cardActions,omitempty"`

	// Header: The header of the card. A header usually contains a title and
	// an image.
	Header *CardHeader `json:"header,omitempty"`

	// Name: Name of the card.
	Name string `json:"name,omitempty"`

	// Sections: Sections are separated by a line divider.
	Sections []*Section `json:"sections,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CardActions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CardActions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// CardAction: A card action is
// the action associated with the card. For an invoice card, a
// typical action would be: delete invoice, email invoice or open
// the
// invoice in browser.
type CardAction struct {
	// ActionLabel: The label used to be displayed in the action menu item.
	ActionLabel string `json:"actionLabel,omitempty"`

	// OnClick: The onclick action for this action item.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

type CardHeader struct {
	// ImageStyle: The image's type (e.g. square border or circular border).
	//
	// Possible values:
	//   "IMAGE_STYLE_UNSPECIFIED"
	//   "IMAGE" - Square border.
	//   "AVATAR" - Circular border.
	ImageStyle string `json:"imageStyle,omitempty"`

	// ImageUrl: The URL of the image in the card header.
	ImageUrl string `json:"imageUrl,omitempty"`

	// Subtitle: The subtitle of the card header.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: The title must be specified. The header has a fixed height: if
	// both a
	// title and subtitle is specified, each will take up 1 line. If only
	// the
	// title is specified, it will take up both lines.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageStyle") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageStyle") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// DeprecatedEvent: Hangouts Chat events.
type DeprecatedEvent struct {
	// Action: The form action data associated with an interactive card that
	// was clicked.
	// Only populated for
	// CARD_CLICKED events.
	// See the [Interactive Cards
	// guide](/hangouts/chat/how-tos/cards-onclick) for
	// more information.
	Action *FormAction `json:"action,omitempty"`

	// ConfigCompleteRedirectUrl: The URL the bot should redirect the user
	// to after they have completed an
	// authorization or configuration flow outside of Hangouts Chat. See
	// the
	// [Authorizing access to 3p services
	// guide](/hangouts/chat/how-tos/auth-3p)
	// for more information.
	ConfigCompleteRedirectUrl string `json:"configCompleteRedirectUrl,omitempty"`

	// EventTime: The timestamp indicating when the event was dispatched.
	EventTime string `json:"eventTime,omitempty"`

	// Message: The message that triggered the event, if applicable.
	Message *Message `json:"message,omitempty"`

	// Space: The room or DM in which the event occurred.
	Space *Space `json:"space,omitempty"`

	// ThreadKey: The bot-defined key for the thread related to the event.
	// See the
	// thread_key field of the
	// `spaces.message.create` request for more information.
	ThreadKey string `json:"threadKey,omitempty"`

	// Token: A secret value that bots can use to verify if a request is
	// from Google. The
	// token is randomly generated by Google, remains static, and can be
	// obtained
	// from the Hangouts Chat API configuration page in the Cloud
	// Console.
	// Developers can revoke/regenerate it if needed from the same page.
	Token string `json:"token,omitempty"`

	// Type: The type of the event.
	//
	// Possible values:
	//   "UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "MESSAGE" - A message was sent in a room or direct message.
	//   "ADDED_TO_SPACE" - The bot was added to a room or DM.
	//   "REMOVED_FROM_SPACE" - The bot was removed from a room or DM.
	//   "CARD_CLICKED" - The bot's interactive card was clicked.
	Type string `json:"type,omitempty"`

	// User: The user that triggered the event.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// FormAction: A form action describes the behavior when the form is
// submitted.
// For example, an Apps Script can be invoked to handle the form.
type FormAction struct {
	// ActionMethodName: Apps Script function to invoke when the containing
	// element is
	// clicked/activated.
	ActionMethodName string `json:"actionMethodName,omitempty"`

	// Parameters: List of action parameters.
	Parameters []*ActionParameter `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionMethodName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionMethodName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

// Image: An image that is specified by a URL and can have an onclick
// action.
type Image struct {
	// AspectRatio: The aspect ratio of this image (width/height).
	AspectRatio float64 `json:"aspectRatio,omitempty"`

	// ImageUrl: The URL of the image.
	ImageUrl string `json:"imageUrl,omitempty"`

	// OnClick: The onclick action.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AspectRatio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AspectRatio") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// ImageButton: An image button with an onclick action.
type ImageButton struct {
	// Icon: The icon specified by an enum that indices to an icon provided
	// by Chat
	// API.
	//
	// Possible values:
	//   "ICON_UNSPECIFIED"
	//   "AIRPLANE"
	//   "BOOKMARK"
	//   "BUS"
	//   "CAR"
	//   "CLOCK"
	//   "CONFIRMATION_NUMBER_ICON"
	//   "DOLLAR"
	//   "DESCRIPTION"
	//   "EMAIL"
	//   "EVENT_PERFORMER"
	//   "EVENT_SEAT"
	//   "FLIGHT_ARRIVAL"
	//   "FLIGHT_DEPARTURE"
	//   "HOTEL"
	//   "HOTEL_ROOM_TYPE"
	//   "INVITE"
	//   "MAP_PIN"
	//   "MEMBERSHIP"
	//   "MULTIPLE_PEOPLE"
	//   "OFFER"
	//   "PERSON"
	//   "PHONE"
	//   "RESTAURANT_ICON"
	//   "SHOPPING_CART"
	//   "STAR"
	//   "STORE"
	//   "TICKET"
	//   "TRAIN"
	//   "VIDEO_CAMERA"
	//   "VIDEO_PLAY"
	Icon string `json:"icon,omitempty"`

	// IconUrl: The icon specified by a URL.
	IconUrl string `json:"iconUrl,omitempty"`

	// Name: The name of this image_button which will be used for
	// accessibility.
	// Default value will be provided if developers don't specify.
	Name string `json:"name,omitempty"`

	// OnClick: The onclick action.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Icon") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Icon") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// KeyValue: A UI element contains a key (label) and a value (content).
// And this
// element may also contain some actions such as onclick button.
type KeyValue struct {
	// BottomLabel: The text of the bottom label. Formatted text supported.
	BottomLabel string `json:"bottomLabel,omitempty"`

	// Button: A button that can be clicked to trigger an action.
	Button *Button `json:"button,omitempty"`

	// Content: The text of the content. Formatted text supported and always
	// required.
	Content string `json:"content,omitempty"`

	// ContentMultiline: If the content should be multiline.
	ContentMultiline bool `json:"contentMultiline,omitempty"`

	// Icon: An enum value that will be replaced by the Chat API with
	// the
	// corresponding icon image.
	//
	// Possible values:
	//   "ICON_UNSPECIFIED"
	//   "AIRPLANE"
	//   "BOOKMARK"
	//   "BUS"
	//   "CAR"
	//   "CLOCK"
	//   "CONFIRMATION_NUMBER_ICON"
	//   "DOLLAR"
	//   "DESCRIPTION"
	//   "EMAIL"
	//   "EVENT_PERFORMER"
	//   "EVENT_SEAT"
	//   "FLIGHT_ARRIVAL"
	//   "FLIGHT_DEPARTURE"
	//   "HOTEL"
	//   "HOTEL_ROOM_TYPE"
	//   "INVITE"
	//   "MAP_PIN"
	//   "MEMBERSHIP"
	//   "MULTIPLE_PEOPLE"
	//   "OFFER"
	//   "PERSON"
	//   "PHONE"
	//   "RESTAURANT_ICON"
	//   "SHOPPING_CART"
	//   "STAR"
	//   "STORE"
	//   "TICKET"
	//   "TRAIN"
	//   "VIDEO_CAMERA"
	//   "VIDEO_PLAY"
	Icon string `json:"icon,omitempty"`

	// IconUrl: The icon specified by a URL.
	IconUrl string `json:"iconUrl,omitempty"`

	// OnClick: The onclick action. Only the top label, bottom label and
	// content region
	// are clickable.
	OnClick *OnClick `json:"onClick,omitempty"`

	// TopLabel: The text of the top label. Formatted text supported.
	TopLabel string `json:"topLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BottomLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BottomLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

type ListMembershipsResponse struct {
	// Memberships: List of memberships in the requested (or first) page.
	Memberships []*Membership `json:"memberships,omitempty"`

	// NextPageToken: Continuation token to retrieve the next page of
	// results. It will be empty
	// for the last page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Memberships") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Memberships") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

type ListSpacesResponse struct {
	// NextPageToken: Continuation token to retrieve the next page of
	// results. It will be empty
	// for the last page of results. Tokens expire in an hour. An error is
	// thrown
	// if an expired token is passed.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Spaces: List of spaces in the requested (or first) page.
	Spaces []*Space `json:"spaces,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Membership: Represents a membership relation in Hangouts Chat.
type Membership struct {
	// CreateTime: The creation time of the membership a.k.a the time at
	// which the member
	// joined the space, if applicable.
	CreateTime string `json:"createTime,omitempty"`

	// Member: Member details.
	Member *User `json:"member,omitempty"`

	// Name: Resource name of the membership, in the form
	// "spaces/*/members/*".
	//
	// Example: spaces/AAAAMpdlehY/members/105115627578887013105
	Name string `json:"name,omitempty"`

	// State: State of the membership.
	//
	// Possible values:
	//   "MEMBERSHIP_STATE_UNSPECIFIED" - Default, do not use.
	//   "JOINED" - The user has joined the space.
	//   "INVITED" - The user has been invited, is able to join the space,
	// but currently has
	// not joined.
	//   "NOT_A_MEMBER" - The user is not a member of the space, has not
	// been invited and is not
	// able to join the space.
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Message: A message in Hangouts Chat.
type Message struct {
	// ActionResponse: Input only. Parameters that a bot can use to
	// configure how its response is
	// posted.
	ActionResponse *ActionResponse `json:"actionResponse,omitempty"`

	// Annotations: Output only. Annotations associated with the text in
	// this message.
	Annotations []*Annotation `json:"annotations,omitempty"`

	// ArgumentText: Plain-text body of the message with all bot mentions
	// stripped out.
	ArgumentText string `json:"argumentText,omitempty"`

	// Cards: Rich, formatted and interactive cards that can be used to
	// display UI
	// elements such as: formatted texts, buttons, clickable images. Cards
	// are
	// normally displayed below the plain-text body of the message.
	Cards []*Card `json:"cards,omitempty"`

	// CreateTime: Output only. The time at which the message was created in
	// Hangouts Chat
	// server.
	CreateTime string `json:"createTime,omitempty"`

	// FallbackText: A plain-text description of the message's cards, used
	// when the actual cards
	// cannot be displayed (e.g. mobile notifications).
	FallbackText string `json:"fallbackText,omitempty"`

	// Name: Resource name, in the form "spaces/*/messages/*".
	//
	// Example: spaces/AAAAMpdlehY/messages/UMxbHmzDlr4.UMxbHmzDlr4
	Name string `json:"name,omitempty"`

	// PreviewText: Text for generating preview chips. This text will not be
	// displayed to the
	// user, but any links to images, web pages, videos, etc. included here
	// will
	// generate preview chips.
	PreviewText string `json:"previewText,omitempty"`

	// Sender: The user who created the message.
	Sender *User `json:"sender,omitempty"`

	// Space: The space the message belongs to.
	Space *Space `json:"space,omitempty"`

	// Text: Plain-text body of the message.
	Text string `json:"text,omitempty"`

	// Thread: The thread the message belongs to.
	Thread *Thread `json:"thread,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionResponse") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionResponse") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

// OnClick: An onclick action (e.g. open a link).
type OnClick struct {
	// Action: A form action will be trigger by this onclick if specified.
	Action *FormAction `json:"action,omitempty"`

	// OpenLink: This onclick triggers an open link action if specified.
	OpenLink *OpenLink `json:"openLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// OpenLink: A link that opens a new window.
type OpenLink struct {
	// Url: The URL to open.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Url") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Section: A section contains a collection of widgets that are
// rendered
// (vertically) in the order that they are specified. Across all
// platforms,
// cards have a narrow fixed width, so
// there is currently no need for layout properties (e.g. float).
type Section struct {
	// Header: The header of the section, text formatted supported.
	Header string `json:"header,omitempty"`

	// Widgets: A section must contain at least 1 widget.
	Widgets []*WidgetMarkup `json:"widgets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Header") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Header") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Space: A room or DM in Hangouts Chat.
type Space struct {
	// DisplayName: Output only. The display name (only if the space is a
	// room).
	DisplayName string `json:"displayName,omitempty"`

	// Name: Resource name of the space, in the form "spaces/*".
	//
	// Example: spaces/AAAAMpdlehYs
	Name string `json:"name,omitempty"`

	// Type: Output only. The type of a space.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "ROOM" - A chat space where memberships are free to change.
	// Messages in rooms are
	// threaded.
	//   "DM" - 1:1 Direct Message between a human and a bot, where all
	// messages are
	// flat.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// TextButton: A button with text and onclick action.
type TextButton struct {
	// OnClick: The onclick action of the button.
	OnClick *OnClick `json:"onClick,omitempty"`

	// Text: The text of the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OnClick") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OnClick") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// TextParagraph: A paragraph of text. Formatted text supported.
type TextParagraph struct {
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// Thread: A thread in Hangouts Chat.
type Thread struct {
	// Name: Resource name, in the form "spaces/*/threads/*".
	//
	// Example: spaces/AAAAMpdlehY/threads/UMxbHmzDlr4
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// User: A user in Hangouts Chat.
type User struct {
	// DisplayName: The user's display name.
	DisplayName string `json:"displayName,omitempty"`

	// Name: Resource name, in the format "users/*".
	Name string `json:"name,omitempty"`

	// Type: User type.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "HUMAN" - Human user.
	//   "BOT" - Bot user.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// UserMentionMetadata: Annotation metadata for user mentions (@).
type UserMentionMetadata struct {
	// Type: The type of user mention.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "ADD" - Add user to space.
	//   "MENTION" - Mention user in space.
	Type string `json:"type,omitempty"`

	// User: The user mentioned.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

// WidgetMarkup: A widget is a UI element that presents texts, images,
// etc.
type WidgetMarkup struct {
	// Buttons: A list of buttons. Buttons is also oneof data and only one
	// of these
	// fields should be set.
	Buttons []*Button `json:"buttons,omitempty"`

	// Image: Display an image in this widget.
	Image *Image `json:"image,omitempty"`

	// KeyValue: Display a key value item in this widget.
	KeyValue *KeyValue `json:"keyValue,omitempty"`

	// TextParagraph: Display a text paragraph in this widget.
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}
