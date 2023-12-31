import beapi from '@berty/api'

import Basic from './Basic'
import ContactRequestReceived from './ContactRequestReceived'
import ContactRequestSent from './ContactRequestSent'
import GroupInvitation from './GroupInvitation'
import MessageReceived from './MessageReceived'

const T = beapi.messenger.StreamEvent.Notified.Type

const notifications = {
	[T.Unknown]: Basic,
	[T.TypeBasic]: Basic,
	[T.TypeGroupInvitation]: GroupInvitation,
	[T.TypeContactRequestReceived]: ContactRequestReceived,
	[T.TypeContactRequestSent]: ContactRequestSent,
	[T.TypeMessageReceived]: MessageReceived,
}

export default notifications

export { Basic as DefaultNotification }
