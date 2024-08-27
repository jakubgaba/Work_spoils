<script>
	import { fade, slide } from 'svelte/transition';
	import { tests } from '../../connections/connections.ts';
</script>

<div>
	<div class="container-fluid">
		<div>
			{tests}
		</div>
	</div>
</div>
