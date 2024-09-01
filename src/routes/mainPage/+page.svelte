<script>
	import { fade, slide } from 'svelte/transition';
	import { tests } from '../../connections/connections.ts';
</script>

<div>
    <div class="container-fluid">
        {#if $tests.length > 0}
            {#each $tests as test}
                <div class="test-item" transition:fade>
                    <p>ID: {test.id}</p>
                    <p>Name: {test.name}</p>
                    <p>Location: {test.location}</p>
                </div>
            {/each}
        {:else}
            <p>No test data available.</p>
        {/if}
    </div>
</div>