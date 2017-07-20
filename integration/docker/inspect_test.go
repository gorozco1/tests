// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"fmt"

	. "github.com/clearcontainers/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func inspectFormatOptions (formatOption string) TableEntry {
	return Entry(fmt.Sprintf("inspect with %s", formatOption), formatOption)
}

var _= Describe ("inspect", func() {
	var (
		args []string
		id   string
	)

	BeforeEach(func() {
		id = RandID(30)
		args = []string{"run", "-t", "--name", id, Image, "true"}
		command := NewCommand(Docker, args...)
		Expect(command).ToNot(BeNil())
		exitCode := command.Run()
		Expect(exitCode).To(Equal(0))

	})

	AfterEach(func() {
		Expect(ContainerRemove(id)).To(BeTrue())
		Expect(ContainerExists(id)).NotTo(BeTrue())
	})

	DescribeTable("inspect with docker",
		func(formatOption string) {
			args = []string{"inspect", id, "--format"}
			args = append(args, formatOption)

			command := NewCommand(Docker, args...)
			exitCode := command.Run()
			Expect(exitCode).To(Equal(0))
			Expect(command.Stdout.String()).To(ContainSubstring(Image))
 			Expect(command.Stdout.String()).NotTo(BeEmpty())
		},
		inspectFormatOptions("'{{.Config.Image}}'"),
		inspectFormatOptions("'{{json .Config}}'"),
	)
})
