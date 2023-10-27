import * as THREE from 'three';

const canvasElement = document.getElementById('canvas') as HTMLElement;

if (canvasElement) {  // Check if on art page
    const scene: THREE.Scene = new THREE.Scene();
    const camera: THREE.PerspectiveCamera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
    const renderer: THREE.WebGLRenderer = new THREE.WebGLRenderer({ antialias: true });

    renderer.setSize(window.innerWidth, window.innerHeight);
    canvasElement.appendChild(renderer.domElement);

    const geometry: THREE.BoxGeometry = new THREE.BoxGeometry();
    const material: THREE.MeshBasicMaterial = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
    const cube: THREE.Mesh = new THREE.Mesh(geometry, material);
    scene.add(cube);

    camera.position.z = 5;

    const animate = (): void => {
        requestAnimationFrame(animate);
        cube.rotation.x += 0.01;
        cube.rotation.y += 0.01;
        renderer.render(scene, camera);
    };

    animate();
}
