'use client';

import { useState } from 'react';
import { PositionedNumber } from '@/lib/gameUtils';

type Props = {
  numbers: PositionedNumber[];
  startTime: number | null;
  onClear: (end: number) => void;
};

export default function PlayScreen({ numbers, startTime, onClear }: Props) {
  const [nextIndex, setNextIndex] = useState(0);
  const [wrongIndex, setWrongIndex] = useState<number | null>(null);

  const handleClick = (num: number, index: number) => {
    const expected = numbers[nextIndex]?.value;
    if (num === expected) {
      if (nextIndex === numbers.length - 1 && startTime) {
        onClear(Date.now());
      }
      setNextIndex((prev) => prev + 1);
    } else {
      setWrongIndex(index);
      setTimeout(() => setWrongIndex(null), 500);
    }
  };

  return (
    <>
      {numbers.map((item, index) => {
        const isWrong = wrongIndex === index;
        const isPressed = index < nextIndex;

        return (
          <button
            key={`${item.value}-${index}`}
            onClick={() => handleClick(item.value, index)}
            className={`
              absolute w-20 h-20 rounded-full flex items-center justify-center
              text-white text-2xl font-bold shadow-lg transition-all duration-500 ease-in-out
              ${isPressed
                ? 'opacity-0 pointer-events-none'
                : isWrong
                ? 'bg-red-400 animate-pulse'
                : 'bg-blue-400 hover:bg-blue-500'
              }`}
            style={{
              left: `${item.x}%`,
              top: `${item.y}%`,
            }}
          >
            {item.value}
          </button>
        );
      })}
    </>
  );
}
