#!/usr/bin/env node
import * as path from 'path'
import * as childProcess from 'child_process'

const BINARY_DISTRIBUTION_PACKAGES: any = {
  win32_ia32: "calculator-mcp_windows_386_sse2",
  win32_x64: "calculator-mcp_windows_amd64_v1",
  win32_arm64: "calculator-mcp_windows_arm64_v8.0",
  darwin_x64: "calculator-mcp_darwin_amd64_v1",
  darwin_arm64: "calculator-mcp_darwin_arm64_v8.0",
  linux_ia32: "calculator-mcp_linux_386_sse2",
  linux_x64: "calculator-mcp_linux_amd64_v1",
  linux_arm64: "calculator-mcp_linux_arm64_v8.0",
}

function getBinaryPath(): string {
  const suffix = process.platform === 'win32' ? '.exe' : '';
  const pkg = BINARY_DISTRIBUTION_PACKAGES[`${process.platform}_${process.arch}`];
  if (pkg) {
    return path.resolve(__dirname, pkg, `calculator-mcp${suffix}`);
  } else {
    throw new Error(`Unsupported platform: ${process.platform}_${process.arch}`);
  }
}

childProcess.execFileSync(getBinaryPath(), process.argv, {
  stdio: 'inherit',
});
