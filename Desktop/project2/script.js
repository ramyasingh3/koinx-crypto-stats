// Global variables
let array = [];
let arraySize = 50;
let speed = 50;
let isSorting = false;
let currentAlgorithm = null;

// Initialize the application
function init() {
    generateNewArray();
    updateArraySize();
    updateSpeed();
}

// Generate a new random array with VIBGYOR pattern
function generateNewArray() {
    if (isSorting) return;
    
    array = [];
    const colors = [
        { name: 'Violet', value: 300 },
        { name: 'Indigo', value: 250 },
        { name: 'Blue', value: 200 },
        { name: 'Green', value: 150 },
        { name: 'Yellow', value: 100 },
        { name: 'Orange', value: 50 },
        { name: 'Red', value: 0 }
    ];
    
    // Create array with random distribution of colors
    for (let i = 0; i < arraySize; i++) {
        const randomColor = colors[Math.floor(Math.random() * colors.length)];
        array.push({
            value: randomColor.value + Math.random() * 20, // Add some variation
            color: randomColor.name
        });
    }
    displayArray();
}

// Display the array as bars
function displayArray(comparingIndices = [], sortedIndices = []) {
    const container = document.getElementById('array-container');
    container.innerHTML = '';
    
    array.forEach((item, index) => {
        const bar = document.createElement('div');
        bar.className = 'array-bar';
        
        // Calculate height based on position for rainbow pattern
        if (sortedIndices.length > 0) {
            const position = sortedIndices.indexOf(index);
            if (position !== -1) {
                // Create rainbow pattern
                const maxHeight = 300;
                const heightRatio = (position + 1) / array.length;
                bar.style.height = `${maxHeight * heightRatio}px`;
                
                // Add individual rainbow color based on position
                const colors = [
                    '#9400D3', // Violet
                    '#4B0082', // Indigo
                    '#0000FF', // Blue
                    '#00FF00', // Green
                    '#FFFF00', // Yellow
                    '#FFA500', // Orange
                    '#FF0000'  // Red
                ];
                
                // Calculate which color to use based on position
                const colorIndex = Math.floor((position / array.length) * colors.length);
                bar.style.background = colors[colorIndex];
                bar.style.backgroundSize = '100% 100%';
                bar.style.backgroundPosition = '0 0';
            } else {
                bar.style.height = `${item.value}px`;
                bar.style.background = '#3498db'; // Default blue color for unsorted bars
            }
        } else {
            bar.style.height = `${item.value}px`;
            bar.style.background = '#3498db'; // Default blue color for unsorted bars
        }
        
        if (comparingIndices.includes(index)) {
            bar.classList.add('comparing');
            bar.style.background = '#e74c3c'; // Red color for comparing
        } else if (sortedIndices.includes(index)) {
            bar.classList.add('sorted');
        }
        
        container.appendChild(bar);
    });
}

// Update array size
function updateArraySize() {
    arraySize = parseInt(document.getElementById('arraySize').value);
    document.getElementById('arraySizeValue').textContent = arraySize;
    generateNewArray();
}

// Update sorting speed
function updateSpeed() {
    speed = parseInt(document.getElementById('speed').value);
    document.getElementById('speedValue').textContent = speed;
}

// Sleep function for animation
const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

// Start sorting with selected algorithm
async function startSort(algorithm) {
    if (isSorting) return;
    
    isSorting = true;
    currentAlgorithm = algorithm;
    
    switch (algorithm) {
        case 'bubble':
            await bubbleSort();
            break;
        case 'insertion':
            await insertionSort();
            break;
        case 'merge':
            await mergeSort();
            break;
        case 'quick':
            await quickSort();
            break;
    }
    
    isSorting = false;
    currentAlgorithm = null;
}

// Bubble Sort
async function bubbleSort() {
    const n = array.length;
    for (let i = 0; i < n - 1; i++) {
        for (let j = 0; j < n - i - 1; j++) {
            displayArray([j, j + 1]);
            await sleep(speed);
            
            if (array[j].value > array[j + 1].value) {
                [array[j], array[j + 1]] = [array[j + 1], array[j]];
            }
        }
        displayArray([], Array.from({length: n - i}, (_, k) => n - 1 - k));
    }
    displayArray([], Array.from({length: n}, (_, i) => i));
}

// Insertion Sort
async function insertionSort() {
    const n = array.length;
    for (let i = 1; i < n; i++) {
        let key = array[i];
        let j = i - 1;
        
        while (j >= 0 && array[j].value > key.value) {
            displayArray([j, j + 1]);
            await sleep(speed);
            
            array[j + 1] = array[j];
            j--;
        }
        array[j + 1] = key;
        displayArray([], Array.from({length: i + 1}, (_, k) => k));
    }
    displayArray([], Array.from({length: n}, (_, i) => i));
}

// Merge Sort
async function mergeSort() {
    await mergeSortHelper(0, array.length - 1);
    displayArray([], Array.from({length: array.length}, (_, i) => i));
}

async function mergeSortHelper(left, right) {
    if (left < right) {
        const mid = Math.floor((left + right) / 2);
        await mergeSortHelper(left, mid);
        await mergeSortHelper(mid + 1, right);
        await merge(left, mid, right);
    }
}

async function merge(left, mid, right) {
    const n1 = mid - left + 1;
    const n2 = right - mid;
    
    const leftArray = array.slice(left, mid + 1);
    const rightArray = array.slice(mid + 1, right + 1);
    
    let i = 0, j = 0, k = left;
    
    while (i < n1 && j < n2) {
        displayArray([k]);
        await sleep(speed);
        
        if (leftArray[i].value <= rightArray[j].value) {
            array[k] = leftArray[i];
            i++;
        } else {
            array[k] = rightArray[j];
            j++;
        }
        k++;
    }
    
    while (i < n1) {
        displayArray([k]);
        await sleep(speed);
        array[k] = leftArray[i];
        i++;
        k++;
    }
    
    while (j < n2) {
        displayArray([k]);
        await sleep(speed);
        array[k] = rightArray[j];
        j++;
        k++;
    }
    
    displayArray([], Array.from({length: right - left + 1}, (_, i) => left + i));
}

// Quick Sort
async function quickSort() {
    await quickSortHelper(0, array.length - 1);
    displayArray([], Array.from({length: array.length}, (_, i) => i));
}

async function quickSortHelper(low, high) {
    if (low < high) {
        const pivotIndex = await partition(low, high);
        await quickSortHelper(low, pivotIndex - 1);
        await quickSortHelper(pivotIndex + 1, high);
    }
}

async function partition(low, high) {
    const pivot = array[high];
    let i = low - 1;
    
    for (let j = low; j < high; j++) {
        displayArray([j, high]);
        await sleep(speed);
        
        if (array[j].value < pivot.value) {
            i++;
            [array[i], array[j]] = [array[j], array[i]];
        }
    }
    
    displayArray([i + 1, high]);
    await sleep(speed);
    [array[i + 1], array[high]] = [array[high], array[i + 1]];
    
    displayArray([], Array.from({length: high - i}, (_, k) => i + 1 + k));
    return i + 1;
}

// Initialize the application when the page loads
window.onload = init; 