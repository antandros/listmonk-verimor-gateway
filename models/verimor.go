// Copyright (C) 2023 Antandros Teknoloji A.Ş.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package models

type Verimor struct {
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	SourceAddr string `json:"source_addr,omitempty"`
	ValidFor   string `json:"valid_for,omitempty"`
	SendAt     string `json:"send_at,omitempty"`
	CustomID   string `json:"custom_id,omitempty"`
	Datacoding string `json:"datacoding,omitempty"`
	Messages   []SMS  `json:"messages,omitempty"`
}

type SMS struct {
	Msg  string `json:"msg,omitempty"`
	Dest string `json:"dest,omitempty"`
	ID   string `json:"id,omitempty"`
}
