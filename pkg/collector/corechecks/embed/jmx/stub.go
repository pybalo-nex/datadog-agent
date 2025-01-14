// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !jmx

package jmx

import (
	"github.com/DataDog/datadog-agent/comp/agent/jmxlogger"
	dogstatsdServer "github.com/DataDog/datadog-agent/comp/dogstatsd/server"
)

// InitRunner is a stub for builds that do not include jmx
func InitRunner(_ dogstatsdServer.Component, _ jmxlogger.Component) {}
