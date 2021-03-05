<?php
namespace Test;

use Library\UserHandle;
use PHPUnit\Framework\TestCase;

require __DIR__ . '/../src/library/UserHandle.php';

class UserHandleTest extends TestCase
{

    protected $object;

    protected function setUp(): void
    {
        $this->object = new UserHandle();
    }

    /**
     * @covers Library\UserHandle::handle
     * @todo   Implement testHandle().
     */
    public function testHandle(array $data): void
    {
        // 在类UserHandle下模拟一个getUsersByJson的方法,并设定返回值
        // $mock = $this->getMockBuilder(UserHandle::class)->setMethods(['getUsersByJson'])->getMock();
        // $mock->expects($this->once())->method('getUsersByJson')->willReturn(reset($data));
        // $logs = $mock->handle();

        // $this->assertStringStartsWith(key($data), $logs[0]);


        $s = $data;
        $this->assertStringStartsWith('a', 'aa');
    }



}
