// app/game/LevelSelect.tsx
'use client';

import { LevelOption } from '../../../lib/gameUtils';

type Props = {
  onSelectLevel: (level: LevelOption) => void;
};

export default function LevelSelect({ onSelectLevel }: Props) {
  return (
    <div className="flex flex-col items-center mt-10 space-y-4">
      <p className="text-lg font-medium">レベルを選んでね！</p>
      {(['Lv1', 'Lv2', 'Lv3', 'Lv4'] as LevelOption[]).map((lv) => (
        <button
          key={lv}
          onClick={() => onSelectLevel(lv)}
          className="bg-yellow-500 text-white px-6 py-2 rounded hover:bg-yellow-600"
        >
          {lv}
        </button>
      ))}
    </div>
  );
}
