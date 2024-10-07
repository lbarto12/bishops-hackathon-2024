<script>
    import { onMount, onDestroy } from "svelte";
    import { initializeChart, stopPolling } from '../chartUtility';

    let canvasElement;

    onMount(() => {
        initializeChart('pie', canvasElement);
    });

    onDestroy(() => {
        if (stopPolling) stopPolling();
    });
</script>

<canvas bind:this={canvasElement}></canvas>
