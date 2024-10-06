import Chart from "chart.js/auto";
import { getColorSets } from './colors.js';
import { startPolling, pollsData } from './api/polls';

const { backgroundColors, borderColors } = getColorSets();

export let chartObjects = []; // To be able to multiple chart instances
export let stopPolling;

pollsData.subscribe(data => {
    const pollResults = data.map(({ percentage }) => percentage); // Map to percentages directly
    chartObjects.forEach(chartObject => {
        chartObject.data.datasets[0].data = pollResults; // Updates dataset
        chartObject.update(); // Refresh the chart with new data
    });
});

export function initializeChart(chartType, canvasElement) {
    const chartConfig = {
        type: chartType,
        data: {
            labels: ["Candidate 1", "Candidate 2", "Candidate 3"],
            datasets: [{
                label: chartType === 'bar' ? 'Votes' : '',  // Only apply label for bar charts
                data: [],
                backgroundColor: backgroundColors,
                borderColor: borderColors,
                borderWidth: 1
            }]
        },
        options: chartType === 'bar' ? { scales: { y: { beginAtZero: true } } } : {}  // The option is for the bar charts
    };

    const chartObject = new Chart(canvasElement, chartConfig);
    chartObjects.push(chartObject); // Store the chart instance

    stopPolling = startPolling(5000); // Fetch every 5 seconds (test it)
}
