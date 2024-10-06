<script>
    import { onMount, onDestroy } from "svelte";
    import { initializeChart, stopPolling } from '../chartUtility';

    let canvasElement;

    onMount(() => {
        initializeChart('bar', canvasElement); // Initialize Bar Chart
    });

    onDestroy(() => {
        if (stopPolling) stopPolling(); // Stop polling when component is destroyed
    });
</script>

<canvas bind:this={canvasElement}></canvas>


<style>
    canvas {
        width: 30vw;
        height: 30vh; /* not working*/
    }
</style>
