/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package s3

import (
	"github.com/spf13/pflag"

	"kubesphere.io/openpitrix/pkg/utils/reflectutils"
)

// Options contains configuration to access a s3 service
type Options struct {
	Endpoint                string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Region                  string `json:"region,omitempty" yaml:"region,omitempty"`
	DisableSSL              bool   `json:"disableSSL" yaml:"disableSSL"`
	ForcePathStyle          bool   `json:"forcePathStyle" yaml:"forcePathStyle"`
	AccessKeyID             string `json:"accessKeyID,omitempty" yaml:"accessKeyID,omitempty"`
	SecretAccessKey         string `json:"secretAccessKey,omitempty" yaml:"secretAccessKey,omitempty"`
	SessionToken            string `json:"sessionToken,omitempty" yaml:"sessionToken,omitempty"`
	Bucket                  string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
	FilePath                string `json:"filePath" yaml:"filePath"`
	LeaderElectionNamespace string `json:"leaderElectionNamespace" yaml:"leaderElectionNamespace"`
	LeaderElectionID        string `json:"leaderElectionID" yaml:"leaderElectionID"`
	Swagger                 bool   `json:"swagger" yaml:"swagger"`
}

// NewS3Options creates a default disabled Options(empty endpoint)
func NewS3Options() *Options {
	return &Options{
		Endpoint:                "",
		Region:                  "",
		DisableSSL:              true,
		ForcePathStyle:          true,
		AccessKeyID:             "",
		SecretAccessKey:         "",
		SessionToken:            "",
		Bucket:                  "",
		FilePath:                "",
		LeaderElectionNamespace: "",
		LeaderElectionID:        "",
		Swagger:                 false,
	}
}

// Validate check options values
func (s *Options) Validate() []error {
	var errors []error

	return errors
}

// ApplyTo overrides options if it's valid, which endpoint is not empty
func (s *Options) ApplyTo(options *Options) {
	if s.Endpoint != "" {
		reflectutils.Override(options, s)
	}
}

// AddFlags add options flags to command line flags,
// if s3-endpoint if left empty, following options will be ignored
func (s *Options) AddFlags(fs *pflag.FlagSet, c *Options) {
	fs.StringVar(&s.Endpoint, "s3-endpoint", c.Endpoint, ""+
		"Endpoint to access to s3 object storage service, if left blank, the following options "+
		"will be ignored.")

	fs.StringVar(&s.Region, "s3-region", c.Region, ""+
		"Region of s3 that will access to, like us-east-1.")

	fs.StringVar(&s.AccessKeyID, "s3-access-key-id", c.AccessKeyID, "access key of s2i s3")

	fs.StringVar(&s.SecretAccessKey, "s3-secret-access-key", c.SecretAccessKey, "secret access key of s2i s3")

	fs.StringVar(&s.SessionToken, "s3-session-token", c.SessionToken, "session token of s2i s3")

	fs.StringVar(&s.Bucket, "s3-bucket", c.Bucket, "bucket name of s2i s3")

	fs.BoolVar(&s.DisableSSL, "s3-disable-SSL", c.DisableSSL, "disable ssl")

	fs.BoolVar(&s.ForcePathStyle, "s3-force-path-style", c.ForcePathStyle, "force path style")

	fs.StringVar(&s.FilePath, "s3-file-path", c.FilePath, "chart file path")
	fs.StringVar(&s.LeaderElectionNamespace, "s3-leader-election-namespace", c.LeaderElectionNamespace, "LeaderElectionNamespace")
	fs.StringVar(&s.LeaderElectionID, "s3-LeaderElectionID", c.LeaderElectionID, "bucket name of s2i s3")
}
