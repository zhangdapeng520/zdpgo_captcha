package captcha

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestCaptcha_GenerateB64s(t *testing.T) {
	type fields struct {
		Driver Driver
		Store  Store
	}

	dDigit := DriverDigit{80, 240, 5, 0.7, 5}
	audioDriver := NewDriverAudio(rand.Intn(5), "en")
	tests := []struct {
		name     string
		fields   fields
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		{"mem-digit", fields{&dDigit, DefaultMemoryStore}, "xxxx", "", false},
		{"mem-audio", fields{audioDriver, DefaultMemoryStore}, "xxxx", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCaptcha(tt.fields.Driver, tt.fields.Store)
			gotId, b64s, err := c.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(b64s)

			a := c.Store.Get(gotId, false)
			if !c.Verify(gotId, a, true) {
				t.Error("false")
			}
		})
	}
}

func TestCaptcha_Verify(t *testing.T) {
	type fields struct {
		Driver Driver
		Store  Store
	}
	type args struct {
		id     string
		answer string
		clear  bool
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantMatch bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Captcha{
				Driver: tt.fields.Driver,
				Store:  tt.fields.Store,
			}
			if gotMatch := c.Verify(tt.args.id, tt.args.answer, tt.args.clear); gotMatch != tt.wantMatch {
				t.Errorf("Captcha.Verify() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func TestNewCaptcha(t *testing.T) {
	type args struct {
		driver Driver
		store  Store
	}
	tests := []struct {
		name string
		args args
		want *Captcha
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCaptcha(tt.args.driver, tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaptcha_Generate(t *testing.T) {
	tests := []struct {
		name     string
		c        *Captcha
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotB64s, err := tt.c.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("Captcha.Generate() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotB64s != tt.wantB64s {
				t.Errorf("Captcha.Generate() gotB64s = %v, want %v", gotB64s, tt.wantB64s)
			}
		})
	}
}
