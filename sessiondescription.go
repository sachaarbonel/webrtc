// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

import (
	"github.com/pion/sdp/v3"
)

// SessionDescription is used to expose local and remote session descriptions.
type SessionDescription struct {
	Type SDPType `json:"type"`
	SDP  string  `json:"sdp"`

	// This will never be initialized by callers, internal use only
	parsed *sdp.SessionDescription
}

// Unmarshal is a helper to deserialize the sdp.
func (sd *SessionDescription) Unmarshal() (*sdp.SessionDescription, error) {
	sd.parsed = &sdp.SessionDescription{}
	err := sd.parsed.UnmarshalString(sd.SDP)

	return sd.parsed, err
}

// SupportsICETrickle returns true if the SDP contains the "a=ice-options:trickle" line,
// indicating support for ICE trickle as per RFC 8838.
func (sd *SessionDescription) SupportsICETrickle() bool {
	if sd.parsed == nil {
		_, err := sd.Unmarshal()
		if err != nil {
			return false
		}
	}

	// Check for trickle in the session-level ice-options
	for _, attr := range sd.parsed.Attributes {
		if attr.Key == "ice-options" && attr.Value == "trickle" {
			return true
		}
	}

	// Also check in each media description, as it could be there as well
	for _, mediaDesc := range sd.parsed.MediaDescriptions {
		for _, attr := range mediaDesc.Attributes {
			if attr.Key == "ice-options" && attr.Value == "trickle" {
				return true
			}
		}
	}

	return false
}
