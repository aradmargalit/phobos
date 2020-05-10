import { GithubOutlined } from '@ant-design/icons';
import { Button, Tooltip } from 'antd';
import React from 'react';

const githubURL = 'https://github.com/aradmargalit/phobos';

export default function GitHubLink() {
  return (
    <Tooltip title="Contribute on Github" placement="bottomLeft">
      <Button
        shape="circle"
        icon={<GithubOutlined />}
        href={githubURL}
        target="_blank"
        rel="noopener noreferrer"
      />
    </Tooltip>
  );
}
