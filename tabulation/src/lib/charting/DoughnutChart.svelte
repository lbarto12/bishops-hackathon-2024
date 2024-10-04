<script>
    import { onMount, onDestroy } from "svelte";
    import { initializeChart, stopPolling } from '../chartUtility';

    let canvasElement;

    onMount(() => {
        initializeChart('doughnut', canvasElement); // To initialize Doughnut Chart
    });

    onDestroy(() => {
        if (stopPolling) stopPolling(); // Stop polling when component is destroyed
    });
</script>

<canvas bind:this={canvasElement}></canvas>

<style>
    canvas {
        width: 30vw;
        height: 30vh;
    }
</style>
