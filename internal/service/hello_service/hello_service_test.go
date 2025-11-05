package hello_service

import (
	"context"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHello_ReturnsGreeting(t *testing.T) {
	// Make selection deterministic
	rand.Seed(42)

	svc := NewHelloService()
	got := svc.SayHello("Alice")

	// Must contain the name and end with an exclamation
	assert.True(t, strings.Contains(got, "Alice"), "expected greeting to contain name")
	assert.True(t, strings.HasSuffix(got, "!"), "expected greeting to end with '!'")

	// Must start with one of the known greetings
	greetings := []string{"Hello", "Bonjour", "Halo", "Hola", "Ciao"}
	found := false
	for _, g := range greetings {
		if strings.HasPrefix(got, g) {
			found = true
			break
		}
	}
	assert.True(t, found, "greeting should start with a known salutation, got: %s", got)
}

func TestSayHelloWithContext_Success(t *testing.T) {
	rand.Seed(42)
	svc := NewHelloService()

	got, err := svc.SayHelloWithContext(context.Background(), "Bob")
	require.NoError(t, err)
	assert.Contains(t, got, "Bob")
}

func TestSayHelloWithContext_ContextCanceled(t *testing.T) {
	svc := NewHelloService()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	got, err := svc.SayHelloWithContext(ctx, "Bob")
	require.Error(t, err)
	assert.Equal(t, context.Canceled, err)
	assert.Equal(t, "", got)
}

func TestSayHelloWithContext_NameError(t *testing.T) {
	svc := NewHelloService()
	got, err := svc.SayHelloWithContext(context.Background(), "error")
	require.Error(t, err)
	assert.Equal(t, context.Canceled, err)
	assert.Equal(t, "", got)
}
