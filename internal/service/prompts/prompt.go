package prompts

func Generate() []string {
	prompts := make([]string, 0, 10)

	//Ask for Provider API documentation
	prompts = append(prompts, "For %[1]s provider give me current link for %[1]s REST API documentation")

	//Ask for how to create account on provider site
	prompts = append(prompts, "For %[1]s provider give me current information about how to create account")

	//Ask how to create provider app
	prompts = append(prompts, "How do I create an app on %[1]s provider site?")

	//Ask what types of authentication does API support
	prompts = append(prompts, "What types of authentication does %[1]s REST API provider support?")

	//Ask in what time access and refresh token expires
	prompts = append(prompts, "In what time access and refresh token expires for %[1]s provider?")

	//Ask for all type of recourses
	prompts = append(prompts, "What are all the resources available in %[1]s provider API?")
	return prompts
}
