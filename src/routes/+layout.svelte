
<script>
	console.log("Iam angry")
</script>

<svelte:head>
	<title>AngryEmployee demo app</title>
</svelte:head>

<nav>
	<a href="/">home</a>
	<a href="/about">about</a>
</nav>

<style>
	:global(body) {
		font-family: 'Open Sans', sans-serif;
	}
</style>

<slot />
