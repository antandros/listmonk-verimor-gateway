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

type Listmonk struct {
	Subject     string       `json:"subject,omitempty"`
	Body        string       `json:"body,omitempty"`
	ContentType string       `json:"content_type,omitempty"`
	Recipients  []Recipients `json:"recipients,omitempty"`
	Campaign    Campaign     `json:"campaign,omitempty"`
}

type Recipients struct {
	UUID    string  `json:"uuid,omitempty"`
	Email   string  `json:"email,omitempty"`
	Name    string  `json:"name,omitempty"`
	Status  string  `json:"status,omitempty"`
	Attribs Attribs `json:"attribs,omitempty"`
}

type Attribs struct {
	Phone   string `json:"phone,omitempty"`
	Company string `json:"company,omitempty"`
}

type Campaign struct {
	UUID string   `json:"uuid,omitempty"`
	Name string   `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
