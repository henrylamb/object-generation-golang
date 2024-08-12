# Guide on Constructing a Definition in Go

When constructing a definition struct in Go, it is important to follow a clear and consistent format. This guide will walk you through how to define a struct, add fields with appropriate types and tags, and provide documentation for each field.

## Example Definitions

To illustrate the process, we'll use modified examples inspired by the original code provided.

### Basic Definition Structure

A definition typically starts with a `struct` declaration. Each field within the struct should have a name, a type, optional JSON tags, and a comment explaining the purpose of the field.

```go
type Definition struct {
    // Type specifies the data type of the schema.
    Type DataType `json:"type,omitempty"`
    
    // Instruction is the instruction for what to generate.
    Instruction string `json:"instruction,omitempty"`
    
    // Properties describes the properties of an object if the schema type is Object.
    Properties map[string]Definition `json:"properties,omitempty"`
    
    // Required specifies which properties are required if the schema type is Object.
    Required []string `json:"required,omitempty"`
    
    // Items specifies which data type an array contains if the schema type is Array.
    Items *Definition `json:"items,omitempty"`
    
    // Model specifies the model type used for generation.
    Model ModelType `json:"model,omitempty"`
    
    // ProcessingOrder defines the order in which fields should be processed.
    ProcessingOrder []string `json:"processingOrder,omitempty"`
    
    // SystemPrompt allows specifying a custom system prompt at the properties level.
    SystemPrompt *string `json:"systemPrompt,omitempty"`
    
    // Req allows sending a request to external services for additional information.
    Req *RequestFormat `json:"req,omitempty"`
    
    // NarrowFocus allows focusing on specific properties for a generation request.
    NarrowFocus *Focus `json:"narrowFocus,omitempty"`
}
```

### Auxiliary Structs

In addition to the main `Definition` struct, you might need auxiliary structs for more specific tasks.

#### Focus Struct

```go
// Focus allows sending a focused request to an LLM with minimal context.
type Focus struct {
    // Prompt is the prompt sent to the LLM.
    Prompt string `json:"prompt"`
    
    // Fields denotes the properties to extract and their order.
    Fields []string `json:"fields"`
}
```

#### RequestFormat Struct

```go
// RequestFormat defines the structure of an external request.
type RequestFormat struct {
    // URL is the endpoint to send the request to.
    URL string `json:"url"`
    
    // Method is the HTTP method to use for the request.
    Method HTTPMethod `json:"method"`
    
    // Headers are optional headers to include in the request.
    Headers map[string]string `json:"headers,omitempty"`
    
    // Body is the optional body content for the request.
    Body map[string]interface{} `json:"body,omitempty"`
    
    // Authorization is the optional authorization token.
    Authorization string `json:"authorization,omitempty"`
}
```

### Example Function Using Definitions

Here is an example function that constructs a book definition using nested definitions for pages and chapters.

```go
func CreateBookDefinition(writingStyle, projectId, systemPrompt string) *Definition {

    // Definition for a page
    pageDefinition := Definition{
        Type:        Object,
        Instruction: "Create the details and information for a single page within a book.",
        Properties: map[string]Definition{
            "outline": {
                Type:        String,
                Instruction: "Create a detailed overview of the content for this page.",
            },
            "finalContent": {
                Type: String,
                Req: &RequestFormat{
                    URL:           "http://localhost:8001/getData",
                    Method:        POST,
                    Body:          map[string]interface{}{"projectId": projectId},
                    Authorization: fmt.Sprintf("Bearer %s", os.Getenv("API_KEY")),
                },
                Instruction: fmt.Sprintf(`
Create a vivid and immersive page for a fiction book based on the outline provided. Focus on detailed descriptions and engaging dialogue. 
Writing Style: %s
Minimum word count: 400 words.
`, writingStyle),
            },
            "editedContent": {
                Type: String,
                NarrowFocus: &Focus{
                    Prompt: "Format the page contents into paragraphs, maintaining all existing content.",
                    Fields: []string{"finalContent"},
                },
                SystemPrompt: &systemPrompt,
            },
        },
    }

    // Definition for a chapter
    chapterDefinition := Definition{
        Type:        Object,
        Instruction: "Create the details of the chapter, ensuring a cohesive story.",
        Properties: map[string]Definition{
            "title": {
                Type:        String,
                Instruction: "Create a chapter title with a maximum of 5 words.",
            },
            "outline": {
                Type:        String,
                Instruction: "Create a chapter outline detailing the events of the chapter.",
            },
            "pages": {
                Type:        Array,
                Items:       &pageDefinition,
                Instruction: "Create a list of pages, each logically leading to the next.",
            },
        },
    }

    // Definition for a book
    bookDefinition := Definition{
        Type:        Object,
        Instruction: "Create an engaging fictional story based on the provided outline.",
        Properties: map[string]Definition{
            "title": {
                Type:        String,
                Instruction: "Return a book title that encapsulates the story in a maximum of 5 words.",
            },
            "chapters": {
                Type:        Array,
                Items:       &chapterDefinition,
                Instruction: "Using the story outline, create a list of chapters detailing the narrative.",
            },
        },
    }

    return &bookDefinition
}
```

## Best Practices

- **Field Naming:** Use camelCase for field names.
- **Types:** Choose appropriate types for each field (e.g., `string`, `int`, `bool`, `map`, `struct`).
- **Tags:** Use JSON tags to specify how fields should be serialized. Use `omitempty` to omit empty fields.
- **Comments:** Provide clear and concise comments for each field to explain its purpose.
- **Reusability:** Break down complex definitions into smaller, reusable structs.
- **Consistent Formatting:** Ensure consistent indentation and spacing for readability.

By following these guidelines, you can create well-structured and easily understandable definition structs in Go.