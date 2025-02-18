package util

import "strings"

// RemoveTaggedContent removes content that is enclosed between tags, including the tags themselves
func RemoveTaggedContent(content, startTag, endTag string) string {
	for {
		startIdx := strings.Index(content, startTag)
		if startIdx == -1 {
			break
		}

		endIdx := strings.Index(content[startIdx:], endTag)
		if endIdx == -1 {
			break
		}
		endIdx += startIdx + len(endTag)

		// Remove the content between and including the tags
		content = content[:startIdx] + content[endIdx:]
	}

	return content
}
