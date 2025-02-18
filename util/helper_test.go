package util

import "testing"

func TestRemoveIgnoredContent(t *testing.T) {
	startTag := "<!-- ss-markdown-ignore start -->"
	endTag := "<!-- ss-markdown-ignore end -->"

	tests := []struct {
		name     string
		input    string
		startTag string
		endTag   string
		expected string
	}{
		{
			name:     "No ignore tags",
			input:    "This is normal content",
			startTag: startTag,
			endTag:   endTag,
			expected: "This is normal content",
		},
		{
			name: "Single ignore block",
			input: `Before
<!-- ss-markdown-ignore start -->
Ignored content
<!-- ss-markdown-ignore end -->
After`,
			startTag: startTag,
			endTag:   endTag,
			expected: "Before\n\nAfter",
		},
		{
			name: "Multiple ignore blocks",
			input: `Start
<!-- ss-markdown-ignore start -->
Ignore 1
<!-- ss-markdown-ignore end -->
Middle
<!-- ss-markdown-ignore start -->
Ignore 2
<!-- ss-markdown-ignore end -->
End`,
			startTag: startTag,
			endTag:   endTag,
			expected: "Start\n\nMiddle\n\nEnd",
		},
		{
			name: "Incomplete tags",
			input: `Start
<!-- ss-markdown-ignore start -->
Incomplete block`,
			startTag: startTag,
			endTag:   endTag,
			expected: `Start
<!-- ss-markdown-ignore start -->
Incomplete block`,
		},
		{
			name: "Custom tags",
			input: `Start
[CUSTOM-START]
Ignore this
[CUSTOM-END]
End`,
			startTag: "[CUSTOM-START]",
			endTag:   "[CUSTOM-END]",
			expected: "Start\n\nEnd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveTaggedContent(tt.input, tt.startTag, tt.endTag)
			if result != tt.expected {
				t.Errorf("RemoveTaggedContent() = %v, want %v", result, tt.expected)
			}
		})
	}
}
