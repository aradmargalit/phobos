import React from 'react';
import { Cell, Pie, PieChart } from 'recharts';

export default function RadialActivityTypesGraph({ colors, typeBreakdown }) {
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
            <Cell key={entry.name} fill={colors[index % colors.length]} />
          ))}
        </Pie>
      </PieChart>
    </div>
  );
}
