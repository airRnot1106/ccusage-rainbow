package rainbow

import (
	"time"

	"ccusage-rainbow/internal/domain/entities"
	"ccusage-rainbow/internal/domain/interfaces"
)

// RainbowTextUseCase handles the business logic for displaying animated rainbow text
type RainbowTextUseCase struct {
	asciiRenderer interfaces.ASCIIRenderer
	colorAnimator interfaces.ColorAnimator
	animation     *entities.RainbowAnimation
}

// NewRainbowTextUseCase creates a new RainbowTextUseCase
func NewRainbowTextUseCase(
	asciiRenderer interfaces.ASCIIRenderer,
	colorAnimator interfaces.ColorAnimator,
) *RainbowTextUseCase {
	return &RainbowTextUseCase{
		asciiRenderer: asciiRenderer,
		colorAnimator: colorAnimator,
		animation:     entities.NewRainbowAnimation(100 * time.Millisecond),
	}
}

// RenderAnimatedText renders text with animated rainbow colors (uses medium size)
func (uc *RainbowTextUseCase) RenderAnimatedText(text *entities.Text) (string, error) {
	return uc.RenderAnimatedTextWithSize(text, interfaces.FontSizeMedium)
}

// RenderAnimatedTextWithSize renders text with specified font size and animated rainbow colors
func (uc *RainbowTextUseCase) RenderAnimatedTextWithSize(text *entities.Text, size interfaces.FontSize) (string, error) {
	// Render plain ASCII art with specified size
	plainASCII, err := uc.asciiRenderer.RenderPlainWithSize(text, size)
	if err != nil {
		return "", err
	}

	// Apply rainbow colors with current animation state
	coloredASCII := uc.colorAnimator.ApplyRainbowColors(plainASCII, uc.animation)

	return coloredASCII, nil
}

// GetDisplayWidth calculates the display width of the rendered text (uses medium size)
func (uc *RainbowTextUseCase) GetDisplayWidth(text *entities.Text) (int, error) {
	return uc.GetDisplayWidthWithSize(text, interfaces.FontSizeMedium)
}

// GetDisplayWidthWithSize calculates display width for specific font size
func (uc *RainbowTextUseCase) GetDisplayWidthWithSize(text *entities.Text, size interfaces.FontSize) (int, error) {
	return uc.asciiRenderer.GetDisplayWidthWithSize(text, size)
}

// SelectOptimalFontSize chooses the best font size based on terminal dimensions
func (uc *RainbowTextUseCase) SelectOptimalFontSize(text *entities.Text, terminalWidth, terminalHeight int) (interfaces.FontSize, error) {
	// Try large size first
	largeWidth, err := uc.GetDisplayWidthWithSize(text, interfaces.FontSizeLarge)
	if err == nil && largeWidth <= terminalWidth && terminalHeight >= 15 { // Need at least 15 rows for large (10 + padding)
		return interfaces.FontSizeLarge, nil
	}

	// Try medium size
	mediumWidth, err := uc.GetDisplayWidthWithSize(text, interfaces.FontSizeMedium)
	if err == nil && mediumWidth <= terminalWidth && terminalHeight >= 12 { // Need at least 12 rows for medium (7 + padding)
		return interfaces.FontSizeMedium, nil
	}

	// Fall back to small size
	return interfaces.FontSizeSmall, nil
}

// AdvanceAnimation advances the animation to the next frame
func (uc *RainbowTextUseCase) AdvanceAnimation() {
	uc.animation.NextFrame()
}

// GetAnimationInterval returns the animation interval
func (uc *RainbowTextUseCase) GetAnimationInterval() time.Duration {
	return uc.animation.GetInterval()
}
