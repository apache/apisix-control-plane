/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger = zap.NewNop()

var level zap.AtomicLevel = zap.NewAtomicLevel()

func init() {
	Init()
}

// Level returns the current log level so that we can do extra things according
// to the current mode
func Level() zap.AtomicLevel {
	return level
}

func Init() {
	// TODO use different log level in different environment
	lv := zapcore.InfoLevel
	level.SetLevel(lv)

	var cores []zapcore.Core
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	// default to stderr
	cores = append(cores, zapcore.NewCore(consoleEncoder, os.Stderr, level))
	core := zapcore.NewTee(cores...)
	// add line and stack info (under error log level)
	logger = zap.New(core, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller())
}

// Errorw logs the error occurs with the stacktrace
func Errorw(msg string, kv ...interface{}) {
	logger.Sugar().Errorw(msg, kv...)
}

// Warnw logs what happened unexpectedly but don't need to be handled immediately
func Warnw(msg string, kv ...interface{}) {
	logger.Sugar().Warnw(msg, kv...)
}

// Infow logs what happened normally
func Infow(msg string, kv ...interface{}) {
	logger.Sugar().Infow(msg, kv...)
}

// Debugw logs the details for debug
func Debugw(msg string, kv ...interface{}) {
	logger.Sugar().Debugw(msg, kv...)
}
