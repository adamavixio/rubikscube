import * as THREE from "three";
import { CSS2DObject, CSS2DRenderer } from "three/examples/jsm/renderers/CSS2DRenderer.js";

const top = 0;
const front = 1;
const bottom = 2;
const back = 3;
const left = 4;
const right = 5;

const grey = 0x808080;
const white = 0xffffff;
const blue = 0x0000ff;

//
// Setup
//
const scene = new THREE.Scene();
scene.background = new THREE.Color(grey);

const axesHelper = new THREE.AxesHelper(5);
scene.add(axesHelper);

const camera = new THREE.PerspectiveCamera(
	75,
	window.innerWidth / window.innerHeight,
	0.1,
	1000
);
camera.position.set(0, 0, 10)

const renderer = new THREE.WebGLRenderer();
renderer.setPixelRatio(window.devicePixelRatio);
renderer.setSize(window.innerWidth, window.innerHeight);

const labelRenderer = new CSS2DRenderer();
labelRenderer.setSize(window.innerWidth, window.innerHeight);
labelRenderer.domElement.style.position = 'absolute';
labelRenderer.domElement.style.top = '0px';

//
// Cube
//

const createCube = (size: number) => {
	let cube: number[][][] = [], count = 1
	for (let i = 0; i < 6; i++) {
		cube.push([])
		for (let j = 0; j < size; j++) {
			cube[i].push([])
			for (let k = 0; k < size; k++) {
				cube[i][j].push(count++)
			}
		}
	}
	console.log(cube)
	return cube
}

const createCubeLayout = (cube: number[][][]) => {
	const size = cube[0].length
	createCubeSide(size, front);
};

const createCubeSide = (size: number, face: number) => {
	if (face === front) {
		console.log(cube[face])
		cube[face].forEach((row, i) =>
			row.forEach((value, j) =>
				createCubePiece(size, face, i, j, value, blue)
			)
		);
	}
};

const createCubePiece = (
	size: number,
	face: number,
	row: number,
	column: number,
	value: number,
	color: number
) => {
	const scale = 3
	const length = scale / size
	const x = -scale / 2 + column * length + length / 2
	const y = scale / 2 - row * length - length / 2

	const cubeGeometry = new THREE.BoxGeometry(length, length, length);
	const cubeMaterial = new THREE.MeshBasicMaterial({ color: blue });
	const cube = new THREE.Mesh(cubeGeometry, cubeMaterial);
	scene.add(cube);

	const lineGeometry = new THREE.BufferGeometry();
	const lineMaterial = new THREE.LineBasicMaterial({ color: white });
	const line = new THREE.Line(lineGeometry, lineMaterial);
	scene.add(line)

	const labelDiv = document.createElement('div');
	labelDiv.className = 'label';
	labelDiv.textContent = value.toString();
	const label = new CSS2DObject(labelDiv);
	label.position.set(0, 0, 0);
	cube.add(label);

	if (face === front) {
		cube.position.set(x, y, length)

		const positions = new THREE.Float32BufferAttribute(
			[
				x - length / 2, y + length / 2, 1.5 * length,
				x + length / 2, y + length / 2, 1.5 * length,
				x + length / 2, y - length / 2, 1.5 * length,
				x - length / 2, y - length / 2, 1.5 * length,
				x - length / 2, y + length / 2, 1.5 * length,
			],
			3
		)
		lineGeometry.setAttribute('position', positions);
	}
};

const cube = createCube(3)
createCubeLayout(cube);

//
// Resize
//
window.addEventListener("resize", onWindowResize, false);
function onWindowResize() {
	camera.aspect = window.innerWidth / window.innerHeight;
	camera.updateProjectionMatrix();
	renderer.setSize(window.innerWidth, window.innerHeight);
	labelRenderer.setSize(window.innerWidth, window.innerHeight);
}

//
// Animations
//
const animate = () => {
	requestAnimationFrame(animate);

	renderer.render(scene, camera);
	labelRenderer.render(scene, camera);
};
animate();

export { renderer, labelRenderer };
