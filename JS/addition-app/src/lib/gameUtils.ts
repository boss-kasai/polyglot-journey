// lib/gameUtils.ts
export type PositionedNumber = {
    value: number;
    x: number; // percentage
    y: number;
  };
  
  export type LevelOption = 'Lv1' | 'Lv2' | 'Lv3' | 'Lv4';
  
  const BUTTON_SIZE = 64; // px
  const MAX_TRIES = 100;
  
  export function generateNumbers(level: LevelOption): PositionedNumber[] {
    let values: number[] = [];
  
    switch (level) {
      case 'Lv1':
        values = [...Array(11).keys()]; // 0-10
        break;
      case 'Lv2':
        values = [...Array(21).keys()]; // 0-20
        break;
      case 'Lv3':
        values = shuffle([...Array(31).keys()]).slice(0, 20); // 0-30からランダム20個
        break;
      case 'Lv4': {
        const part1 = sampleRange(0, 100, 6);
        const part2 = sampleRange(101, 200, 7);
        const part3 = sampleRange(201, 300, 7);
        values = shuffle([...part1, ...part2, ...part3]);
        break;
      }
    }
  
    return placeNumbersRandomly(values);
  }
  
  function sampleRange(min: number, max: number, count: number): number[] {
    const range = Array.from({ length: max - min + 1 }, (_, i) => i + min);
    return shuffle(range).slice(0, count);
  }
  
  function shuffle<T>(array: T[]): T[] {
    return [...array].sort(() => Math.random() - 0.5);
  }
  
  function placeNumbersRandomly(values: number[]): PositionedNumber[] {
    const containerWidth = typeof window !== 'undefined' ? window.innerWidth : 800;
    const containerHeight = typeof window !== 'undefined' ? window.innerHeight * 0.8 : 600;
  
    const positioned: PositionedNumber[] = [];
    const existingRects: { x: number; y: number }[] = [];
  
    for (const value of values) {
      let tryCount = 0;
      let valid = false;
      let x = 0;
      let y = 0;
  
      while (!valid && tryCount < MAX_TRIES) {
        x = Math.random() * (containerWidth - BUTTON_SIZE);
        y = Math.random() * (containerHeight - BUTTON_SIZE);
  
        const overlaps = existingRects.some((r) => {
          const dx = r.x - x;
          const dy = r.y - y;
          const distance = Math.sqrt(dx * dx + dy * dy);
          return distance < BUTTON_SIZE + 8;
        });
  
        if (!overlaps) {
          valid = true;
          existingRects.push({ x, y });
        }
  
        tryCount++;
      }
  
      positioned.push({
        value,
        x: (x / containerWidth) * 100,
        y: (y / containerHeight) * 100,
      });
    }
  
    return positioned;
  }
  