import React from 'react';
import { Cell, Pie, PieChart } from 'recharts';

const COLORS = [
  '#29bdd6',
  '#11b6d1',
  '#0fa4bc',
  '#0e92a7',
  '#0c7f92',
  '#095b69',
  '#05373f',
];

export default function RadialActivityTypesGraph({ typeBreakdown }) {
  return (
    <div className="statistics--pie">
      <h3>Activity Type Split</h3>
      <PieChart width={500} height={300}>
        <Pie
          data={typeBreakdown}
          innerRadius={40}
          outerRadius={60}
          fill="#8884d8"
          paddingAngle={1}
          dataKey="portion"
          nameKey="name"
          label={d => d.name}
        >
          {typeBreakdown.map((entry, index) => (
            <Cell key={entry.name} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
      </PieChart>
    </div>
  );
}
