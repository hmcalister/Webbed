<script>
	import { PingServer } from "$lib/API";

	var serverPingResponse = "";
	async function doClick() {
		let {statusCode, message} = await PingServer()
		console.log(message)
		serverPingResponse = message
	}
</script>

<div class="container pt-10 h-full mx-auto flex justify-center items-center">
	<div class="space-y-5">
		<h1 class="h1">Webbed Frontend Template</h1>
		<h2 class="h2">Development</h2>
		<p>When developing, run <code>npm run dev</code> from the <code>frontend</code> directory. This allows for normal frontend development using all the nice Node tools. Running <code>go run . -allowCORS</code> from the root directory allows the frontend to access the Go server, even when running on a different port (for development only). Adding JS functions for API requests to <code>frontend/src/lib</code> will allow for clean, easy development of interactions between frontend and backend.</p>
		<p>Alternatively, install <a class="underline" href="https://github.com/cosmtrek/air">air</a> to allow for hot reloading of Go. Once installed, use the included <code>Makefile</code> by running <code>make dev</code> to start both the npm dev server and hot reloading of the Go backend. Ensure during development you are using the npm server, not the embedded static files!</p>
		<h2 class="h2">Deployment</h2>
		<p>When ready to build, run <code>npm run build</code> in the <code>frontend</code> directory, taking care to note the Svelte adaptor has been changed to <code>@sveltejs/adapter-static</code>. This compiles the frontend page to a static site located at <code>frontend/build</code>. Running <code>go build .</code> will then embed these static files in the binary, allowing for the frontend to be served without also packaging additional folders --- only a single executable needs to be distributed!</p>
		<p>This processed is automated in the included <code>Makefile</code>, which can be employed using <code>make build</code>.</p>
		<h2 class="h2">Example Request</h2>
		<div class="flex gap-5 justify-center">	
			<button class="btn variant-filled" on:click={doClick}>Ping Server</button>
			<button class="btn variant-filled" on:click={() => {serverPingResponse = ""}}>Clear Response</button>
		</div>
		<p class="text-center">Response from server:<br>{serverPingResponse}</p>
	</div>
</div>

<style>
	code {
		background-color: #555C;
		padding: 0em 0.5em 0em 0.5em;
		border-radius: 0.3em;
	}
</style>
