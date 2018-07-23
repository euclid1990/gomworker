<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Workers extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('workers', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name')->unique();
            $table->enum('status', ['started', 'running', 'stopped', 'asked_to_stop'])->default('started');
            $table->float('usage_cpu')->unsigned()->default(0);
            $table->float('usage_memory')->unsigned()->default(0);
            $table->string('queue')->default('default');
            $table->boolean('once')->default(false);
            $table->unsignedInteger('delay')->default(0);
            $table->boolean('force')->default(false);
            $table->unsignedInteger('memory')->default(128);
            $table->unsignedInteger('sleep')->default(3);
            $table->unsignedInteger('timeout')->default(60);
            $table->unsignedInteger('tries')->default(0);
            $table->timestamp('started_at')->nullable();
            $table->timestamps();
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('workers');
    }
}
