# Test Generation Template

Generate table-driven tests for the target package/file.

## Style
- Name: Test_<FuncName>_<Case>
- Arrange–Act–Assert sections separated by comments
- Cover: zero/one/many, error path, boundary values

## Checklist
- [ ] Deterministic inputs (no time.Now() without injection)
- [ ] Edge cases (empty, nil, negative, very large)
- [ ] Meaningful failure messages

## Prompt
"""
Using the above template and our style, generate tests for <package>/<file>. 
Prefer small, readable cases over exhaustive ones. Avoid flakiness.
"""
