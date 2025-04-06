resource mydrbd {
	protocol C;

	on hostname1 {
		device /dev/drbd0;
		disk ec2disk1;
		address privateIp1:7789;
		meta-disk internal;
	}

    on hostname2 {
		device /dev/drbd0;
		disk ec2disk2;
		address privateIp2:7789;
		meta-disk internal;
	}
}
