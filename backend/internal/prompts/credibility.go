package prompts

const CredibilitySystemPrompt = `
You are an analysis system that identifies credibility and risk signals.
You do not determine truth or falsehood.
You return structured observations only.
`

const CredibilityUserPrompt = `
Analyze the following content and return:

1. Emotional tone (neutral, emotional, manipulative)
2. Presence of urgency or fear-based language (yes/no)
3. Whether claims appear opinion-based or factual
4. Up to 3 credibility risk signals (short phrases)

Content:
"""
%s
"""
`
