// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

// OfferAnswerOptions is a base structure which describes the options that
// can be used to control the offer/answer creation process.
type OfferAnswerOptions struct {
	// VoiceActivityDetection allows the application to provide information
	// about whether it wishes voice detection feature to be enabled or disabled.
	VoiceActivityDetection bool

	// ICETrickle indicates whether ICE trickling is supported.
	// When true, the a=ice-options:trickle attribute will be added to the SDP.
	// See RFC 8838 for details: https://datatracker.ietf.org/doc/html/rfc8838
	ICETrickle bool
}

// AnswerOptions structure describes the options used to control the answer
// creation process.
type AnswerOptions struct {
	OfferAnswerOptions
}

// OfferOptions structure describes the options used to control the offer
// creation process.
type OfferOptions struct {
	OfferAnswerOptions

	// ICERestart forces the underlying ice gathering process to be restarted.
	// When this value is true, the generated description will have ICE
	// credentials that are different from the current credentials
	ICERestart bool
}
