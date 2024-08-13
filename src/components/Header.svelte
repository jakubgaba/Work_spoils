<script>
	import { fade, fly } from 'svelte/transition';
	import { quintOut } from 'svelte/easing';

    let count = 1;
	$: displayText1 = count === 0 ? '(No Selection)' : `${count} product(s), `;
	$: displayText2 = count === 0 ? '' : ``;
</script>

<div>
	{#key displayText1}
		<span style="display: inline-block" in:fade|global={{ duration: 500 }}>
			{displayText1}
		</span>
	{/key}
    {#key displayText2}
		<span style="display: inline-block" in:fly|global={{ y: -50, duration: 500, easing: quintOut }}>
			{displayText2}
		</span>
	{/key}
	<a href="/#"> Submit Order </a>
</div>
