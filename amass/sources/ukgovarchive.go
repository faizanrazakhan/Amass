// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package sources

type UKGovArchive struct {
	BaseDataSource
	baseURL string
}

func NewUKGovArchive() DataSource {
	u := &UKGovArchive{baseURL: "http://webarchive.nationalarchives.gov.uk"}

	u.BaseDataSource = *NewBaseDataSource(ARCHIVE, "UK Gov Arch")
	return u
}

func (u *UKGovArchive) Query(domain, sub string) []string {
	if sub == "" {
		return []string{}
	}
	return runArchiveCrawler(u.baseURL, domain, sub, u)
}

func (u *UKGovArchive) Subdomains() bool {
	return true
}