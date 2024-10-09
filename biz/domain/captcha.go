/*
 * Copyright 2023 CoolFunAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package domain

type Captcha interface {
	GetCaptcha() (id, b64s string, err error)
}
