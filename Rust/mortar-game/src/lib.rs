use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub struct Point {
    pub x: f64,
    pub y: f64,
}

#[wasm_bindgen]
pub fn simulate_trajectory(angle_deg: f64, velocity: f64, time_step: f64) -> Vec<Point> {
    let angle_rad = angle_deg.to_radians();
    let velocity_x = velocity * angle_rad.cos();
    let velocity_y = velocity * angle_rad.sin();
    let g = 9.81;

    let mut points = Vec::new();
    let mut t = 0.0;
    loop {
        let x = velocity_x * t;
        let y = velocity_y * t - 0.5 * g * t * t;

        if y < 0.0 {
            break;
        }

        points.push(Point { x, y });
        t += time_step;
    }

    points
}

#[wasm_bindgen]
pub fn calculate_distance(x1: f64, y1: f64, x2: f64, y2: f64) -> f64 {
    ((x2 - x1).powi(2) + (y2 - y1).powi(2)).sqrt()
}
