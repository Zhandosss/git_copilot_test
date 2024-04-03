package prompts

func Generate() []string {
	prompts := make([]string, 0, 20)

	//1.Ask for Provider API documentation +
	prompts = append(prompts, "For %[1]s provider give me current link for %[1]s API reference")

	//2.Ask for how to create account on provider site ?
	prompts = append(prompts, "For %[1]s provider give me information about how to create %[1]s account")

	//3.Ask how to create provider app -+
	prompts = append(prompts, "For %[1]s provider document all the steps to create %[1]s provider app. I just need information about how to get client id and client secret for integration")

	//4.Ask what types of authentication does API support
	prompts = append(prompts, "For %[1]s provider give me information about all authentication methods supported by %[1]s provider API")

	//5.Ask in what time access and refresh token expires -+
	prompts = append(prompts, "For %[1]s provider tell me in what time access and refresh tokens expires in %[1]s provider API?")

	//6.Ask for all type of recourses -+
	prompts = append(prompts, "For %[1]s provider give me all API resources available in %[1]s provider API")

	//7.Ask about rate limits
	prompts = append(prompts, "What are the rate limits for %[1]s provider API? Are rate limits different for different endpoints in %[1]s provider API?")

	//8.Ask about rate limit headers
	prompts = append(prompts, "For %[1]s provider API give me information about what rate limit headers are returned in response")

	//9.Ask about rate limit error response +
	prompts = append(prompts, "What is the error response when rate limit is exceeded in %[1]s provider API?")

	//10.Ask about Batch or Bulk operations
	prompts = append(prompts, "Does %[1]s provider API support batch or bulk operations? What API endpoints support batch or bulk operations in %[1]s provider API?")

	//11.Ask about pagination +-
	prompts = append(prompts, "What type of pagination %[1]s provider API supports. What parameter %[1]s provider API supports for pagination. Give some example of pagination work")

	//12.Ask about error response -+
	prompts = append(prompts, "What is the error response format in %[1]s provider API?")

	//13.Ask about API versions +
	prompts = append(prompts, "For %[1]s developer provider how API is versioned")

	//14.Ask about changelog of API +-
	prompts = append(prompts, "Do the %[1]s developers publish API changelog information? If so tell me where i can find %[1]s changelog and give me a link with %[1]s changelog")
	return prompts

}
