<script>
	import { fade, slide } from 'svelte/transition';
	import Header from "../../components/Header.svelte";
</script>

<div>
<Header/>
<div class="container-fluid">
	<div class="row">
		<div in:fade|global={{duration:500}}>
			TEXT
		</div>
	</div>
	</div>
</div>

